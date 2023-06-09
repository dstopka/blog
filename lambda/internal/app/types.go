package app

import "time"

// Post represents a single post returned by queries.
type Post struct {
	Title         string
	Description   string
	Slug          string
	Content       string
	CoverImageURL string
	Tags          []string
	PublishedTime time.Time
	UpdatedTime   *time.Time
}