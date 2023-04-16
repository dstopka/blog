package app

import (
	"context"
)

const defaultPageSize = 10

// QueryPostsPage defines arguments used for querying posts page.
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
//
// If Cursor is not present in the query, first page of the results is returned.
// Maximal number of posts on the page is controlled by PageSize in the query.
func (h QueryPostsPageHandler) Handle(ctx context.Context, query QueryPostsPage) (*PostsPage, error) {
	query.setMissingArguments()

	if query.Cursor == "" {
		return h.postService.FirstPostsPage(ctx, query.PageSize)
	}
	return h.postService.PostsPageFromCursor(ctx, query.Cursor, query.PageSize)
}

func (q *QueryPostsPage) setMissingArguments() {
	if q.PageSize == 0 {
		q.PageSize = defaultPageSize
	}
}
