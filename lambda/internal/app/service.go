package app

import (
	"context"
)

// PostsService defines functions needed to perform post related operations.
//
//go:generate go run github.com/vektra/mockery/v2 --name PostService
type PostService interface {
	// FindPostBySlug finds and returns post identified by the given slug.
	FindPostBySlug(ctx context.Context, slug string) (*Post, error)
	// FirstPostsPage returns the first page of posts. Returned page will contain
	// at most pageSize results.
	FirstPostsPage(ctx context.Context, pageSize int) (*PostsPage, error)
	// PostsPageFromCursor returns the page of posts that start from the provided cursor.
	// Returned page will contain at most pageSize results.
	PostsPageFromCursor(ctx context.Context, cursor string, pageSize int) (*PostsPage, error)
}
