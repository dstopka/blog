package post

import (
	"errors"
	"time"
)

var (
	// ErrEmptyUUID is returned when uuid is empty.
	ErrEmptyUUID = errors.New("empty post id")
	// ErrEmptyTitle is returned when title is empty.
	ErrEmptyTitle = errors.New("empty post title")
	// ErrEmptySlug is returned when slug is empty.
	ErrEmptySlug = errors.New("empty post slug")
	// ErrEmptyDescription is returned when description is empty.
	ErrEmptyDescription = errors.New("empty post description")
	// ErrZeroPublishedTime is returned when published time is zero.
	ErrZeroPublishedTime = errors.New("zero post published time")
	// ErrZeroCoverImage is returned when cover is nil.
	ErrZeroCoverImage = errors.New("zero post cover image")
)

// Post represents a single blog post object.
type Post struct {
	id            string
	title         string
	description   string
	slug          string
	content       string
	tags          []string
	coverImage    CoverImage
	publishedTime time.Time
	updatedTime   time.Time
}

// New returns a new Page created using provided values.
func New(id, title, slug, description string, publishedTime time.Time, coverImg CoverImage) (*Post, error) {
	if id == "" {
		return nil, ErrEmptyUUID
	}

	if title == "" {
		return nil, ErrEmptyTitle
	}

	if slug == "" {
		return nil, ErrEmptySlug
	}

	if description == "" {
		return nil, ErrEmptyDescription
	}

	if publishedTime.IsZero() {
		return nil, ErrZeroPublishedTime
	}

	if coverImg.IsZero() {
		return nil, ErrZeroCoverImage
	}

	post := &Post{
		id:            id,
		title:         title,
		slug:          slug,
		description:   description,
		publishedTime: publishedTime,
		coverImage:    coverImg,
	}

	return post, nil
}

// ID returns post's id.
func (p Post) ID() string {
	return p.id
}

// Title returns post's title.
func (p Post) Title() string {
	return p.title
}

// Description returns post's description.
func (p Post) Description() string {
	return p.description
}

// Slug returns post's slug.
func (p Post) Slug() string {
	return p.slug
}

// Content returns post's content.
func (p Post) Content() string {
	return p.content
}

// CoverImage returns post's coverImage.
func (p Post) CoverImage() CoverImage {
	return p.coverImage
}

// Tags returns post's tags.
func (p Post) Tags() []string {
	return p.tags
}

// PublishedTime returns post's publishedTime.
func (p Post) PublishedTime() time.Time {
	return p.publishedTime
}

// UpdatedTime returns post's updatedTime.
func (p Post) UpdatedTime() time.Time {
	return p.updatedTime
}

// UpdatedAt sets post's updatedTime to the provided t.
func (p *Post) UpdatedAt(t time.Time) {
	p.updatedTime = t
}

// WasUpdated informs whether the post was updated at some point.
func (p Post) WasUpdated() bool {
	return !p.updatedTime.IsZero()
}

// AddTags adds tags to post ignoring empty tags.
func (p *Post) AddTags(tags ...string) {
	for _, tag := range tags {
		if tag != "" {
			p.tags = append(p.tags, tag)
		}
	}
}

// UpdateContent updates post's content to the provided one.
func (p *Post) UpdateContent(content string) {
	p.content = content
}
