package app

import (
	"context"
)

const defaultPageSize = 10

// QueryPostsPage defines arguments used for querying posts page.
//
// If Cursor is not present, first page of the results is returned.
// Maximal number of posts on the page is controlled by PageSize.
type QueryPostsPage struct {
	Cursor   string
	PageSize int
}

// QueryPostsPageHandler contains all dependencies used to query a posts page.
type QueryPostsPageHandler struct {
	postService PostService
}

// NewQueryPostsPageHandler returns a new QueryPostsPageHandler that uses postService.
func NewQueryPostsPageHandler(postService PostService) QueryPostsPageHandler {
	if postService == nil {
		panic("nil postService")
	}

	return QueryPostsPageHandler{postService}
}

// Handle handles querying posts page.
func (h QueryPostsPageHandler) Handle(ctx context.Context, query QueryPostsPage) (*PostsPage, error) {
	if query.PageSize == 0 {
		query.PageSize = defaultPageSize
	}

	if query.Cursor == "" {
		return h.postService.FirstPostsPage(ctx, query.PageSize)
	}
	return h.postService.PostsPageFromCursor(ctx, query.Cursor, query.PageSize)
}
