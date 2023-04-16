package app

import "time"

// Post represents a single post returned by queries.
type Post struct {
	Title         string
	Description   string
	Slug          string
	Content       string
	Tags          []string
	CoverImageURL string
	PublishedTime time.Time
	EditedTime    time.Time
}

// PostsPage represents a page which contains posts.
type PostsPage struct {
	Posts []Post
	Next  string
}