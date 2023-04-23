package adapters

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dstopka/blog/lambda/internal/app"
	"github.com/dstopka/blog/lambda/internal/domain/post"
	"github.com/dstopka/blog/lambda/internal/notion2md"
	"github.com/jomei/notionapi"
)

const (
	propertyDescription   = "Description"
	propertyUpdatedTime   = "Updated time"
	propertyPublished     = "Published"
	propertyPublishedTime = "Published time"
	propertySlug          = "Slug"
	propertyTags          = "Tags"
	propertyTitle         = "Title"
)

var _ app.PostService = (*Notion)(nil)

// Notion implements functions for communicating with notion APIs
// and adapting data to models used by app.
type Notion struct {
	client          *notionapi.Client
	converter       *notion2md.Converter
	postsDatabaseID string
}

// NewNotion returns a new Notion created using client and postDatabaseID.
func NewNotion(client *notionapi.Client, postDatabaseID string) *Notion {
	if client == nil {
		panic("nil client")
	}
	if postDatabaseID == "" {
		panic("empty postDatabaseID")
	}

	conv, err := notion2md.NewConverter(client.Block)
	if err != nil {
		panic(fmt.Errorf("creating notion2md converter failed: %w", err))
	}

	return &Notion{client: client, postsDatabaseID: postDatabaseID, converter: conv}
}

// FindPostBySlug finds and returns post identified by the given slug.
func (n *Notion) FindPostBySlug(ctx context.Context, slug string) (*app.Post, error) {
	req := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.AndCompoundFilter{
			notionapi.PropertyFilter{
				Property: propertySlug,
				RichText: &notionapi.TextFilterCondition{
					Equals: slug,
				},
			},
			notionapi.PropertyFilter{
				Property: propertyPublished,
				Checkbox: &notionapi.CheckboxFilterCondition{
					Equals: true,
				},
			},
		},
	}

	resp, err := n.client.Database.Query(ctx, notionapi.DatabaseID(n.postsDatabaseID), req)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, &app.PostNotFoundError{Slug: slug}
	}
	notionPage := resp.Results[0]

	post, err := n.notionPageToPost(&notionPage)
	if err != nil {
		return nil, err
	}

	content := &bytes.Buffer{}
	if err = n.converter.ConvertContext(ctx, notionPage.ID.String(), content); err != nil {
		return nil, err
	}

	post.UpdateContent(content.String())

	return n.postToQuery(post), nil
}

// FirstPostsPage returns the first page of posts.
func (n *Notion) FirstPostsPage(ctx context.Context, pageSize int) (*app.PostsPage, error) {
	return n.PostsPageFromCursor(ctx, "", pageSize)
}

// PostsPageFromCursor returns the page of posts that start from the provided cursor.
func (n *Notion) PostsPageFromCursor(ctx context.Context, cursor string, pageSize int) (*app.PostsPage, error) {
	req := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: propertyPublished,
			Checkbox: &notionapi.CheckboxFilterCondition{
				Equals: true,
			},
		},
		Sorts: []notionapi.SortObject{
			{
				Property:  propertyPublishedTime,
				Direction: notionapi.SortOrderDESC,
			},
		},
		PageSize:    pageSize,
		StartCursor: notionapi.Cursor(cursor),
	}

	resp, err := n.client.Database.Query(ctx, notionapi.DatabaseID(n.postsDatabaseID), req)
	if err != nil {
		return nil, err
	}

	posts, err := n.notionPagesToQuery(resp.Results)
	if err != nil {
		return nil, err
	}

	return &app.PostsPage{
		Posts: posts,
		Next:  resp.NextCursor.String(),
	}, nil
}

func (n Notion) notionPagesToQuery(pages []notionapi.Page) ([]app.Post, error) {
	queryPosts := make([]app.Post, 0, len(pages))

	for i := range pages {
		post, err := n.notionPageToPost(&pages[i])
		if err != nil {
			return nil, err
		}
		queryPosts = append(queryPosts, *n.postToQuery(post))
	}

	return queryPosts, nil
}

func (n Notion) notionPageToPost(page *notionapi.Page) (*post.Post, error) {
	properties := page.Properties

	propTitle, ok := properties[propertyTitle]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertyTitle}
	}
	title, err := n.titlePropertyAsString(propTitle)
	if err != nil {
		return nil, err
	}

	propDescription, ok := properties[propertyDescription]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertyDescription}
	}
	description, err := n.richTextPropertyAsString(propDescription)
	if err != nil {
		return nil, err
	}

	propSlug, ok := properties[propertySlug]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertySlug}
	}
	slug, err := n.richTextPropertyAsString(propSlug)
	if err != nil {
		return nil, err
	}

	propPublishedTime, ok := properties[propertyPublishedTime]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertyPublishedTime}
	}
	publishedTime, err := n.datePropertyAsTime(propPublishedTime)
	if err != nil {
		return nil, err
	}

	var coverURL string
	if page.Cover != nil {
		coverURL = page.Cover.GetURL()
	}

	coverImg, err := post.NewCoverImage(coverURL)
	if err != nil {
		return nil, err
	}

	post, err := post.New(page.ID.String(), title, slug, description, publishedTime, coverImg)
	if err != nil {
		return nil, err
	}

	propTags, ok := properties[propertyTags]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertyTags}
	}
	tags, err := n.multiSelectPropertyAsStrings(propTags)
	if err != nil {
		return nil, err
	}

	post.AddTags(tags...)

	propUpdatedTime, ok := properties[propertyUpdatedTime]
	if !ok {
		return nil, &app.PropertyNotFoundError{Property: propertyUpdatedTime}
	}
	updatedTime, err := n.datePropertyAsTime(propUpdatedTime)
	if err != nil {
		return nil, err
	}

	post.UpdatedAt(updatedTime)

	return post, nil
}

func (n Notion) postToQuery(post *post.Post) *app.Post {
	queryPost := &app.Post{
		Title:         post.Title(),
		Slug:          post.Slug(),
		Description:   post.Description(),
		PublishedTime: post.PublishedTime(),
		CoverImageURL: post.CoverImage().URL(),
		Tags:          post.Tags(),
		Content:       post.Content(),
	}

	if post.WasUpdated() {
		updatedTime := post.UpdatedTime()
		queryPost.UpdatedTime = &updatedTime
	}

	return queryPost
}

func (n Notion) multiSelectPropertyAsStrings(prop notionapi.Property) ([]string, error) {
	multiSelectProp, ok := prop.(*notionapi.MultiSelectProperty)
	if !ok {
		return nil, errors.New("not a multi select")
	}

	res := make([]string, 0, len(multiSelectProp.MultiSelect))
	for _, opt := range multiSelectProp.MultiSelect {
		res = append(res, opt.Name)
	}

	return res, nil
}

func (n Notion) richTextPropertyAsString(prop notionapi.Property) (string, error) {
	richTextProp, ok := prop.(*notionapi.RichTextProperty)
	if !ok {
		return "", errors.New("not a rich text")
	}

	return n.richTextToString(richTextProp.RichText), nil
}

func (n Notion) titlePropertyAsString(prop notionapi.Property) (string, error) {
	titleProp, ok := prop.(*notionapi.TitleProperty)
	if !ok {
		return "", errors.New("not a title")
	}

	return n.richTextToString(titleProp.Title), nil
}

func (n Notion) datePropertyAsTime(prop notionapi.Property) (time.Time, error) {
	dateProp, ok := prop.(*notionapi.DateProperty)
	if !ok {
		return time.Time{}, errors.New("not a date")
	}

	if dateProp.Date == nil || dateProp.Date.Start == nil {
		return time.Time{}, nil
	}

	return (time.Time)(*dateProp.Date.Start), nil
}

func (n Notion) richTextToString(rt []notionapi.RichText) string {
	var str string
	for _, t := range rt {
		str += t.PlainText
	}

	return str
}
