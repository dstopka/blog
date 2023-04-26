package app

import (
	"context"
)

// QueryAllPostsHandler contains all dependencies used to query a posts page.
type QueryAllPostsHandler struct {
	postService PostService
}

// NewQueryAllPostsHandler returns a new QueryAlllPostsHandler that uses postService.
func NewQueryAllPostsHandler(postService PostService) QueryAllPostsHandler {
	if postService == nil {
		panic("nil postService")
	}

	return QueryAllPostsHandler{postService}
}

// Handle handles querying all posts.
func (h QueryAllPostsHandler) Handle(ctx context.Context) ([]Post, error) {
	return h.postService.AllPosts(ctx)
}
