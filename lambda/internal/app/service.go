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
	// AllPosts returns all published posts.
	AllPosts(ctx context.Context) ([]Post, error)
}
