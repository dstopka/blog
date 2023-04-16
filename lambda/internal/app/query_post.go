package app

import (
	"context"
)

// QueryPost defines arguments used for querying post.
type QueryPost struct {
	Slug string
}

// QueryPostHandler contains all dependencies used to query a post.
type QueryPostHandler struct {
	postService PostService
}

// NewQueryPostHandler returns a new QueryPostHandler that uses postService.
func NewQueryPostHandler(postService PostService) QueryPostHandler {
	if postService == nil {
		panic("nil postService")
	}

	return QueryPostHandler{postService}
}

// Handle handles querying a post.
func (h QueryPostHandler) Handle(ctx context.Context, query QueryPost) (*Post, error) {
	return h.postService.FindPostBySlug(ctx, query.Slug)
}
