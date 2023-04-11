package post

import "errors"

// ErrEmptyCoverImageURL is returned when url is empty.
var ErrEmptyCoverImageURL = errors.New("empty cover image URL")

// CoverImage represents post's cover image.
type CoverImage struct {
	url string
}

// NewCoverImage returns a new CoverImage with the url.
func NewCoverImage(url string) (CoverImage, error) {
	if url == "" {
		return CoverImage{}, ErrEmptyCoverImageURL
	}

	return CoverImage{
		url: url,
	}, nil
}

// IsZero returns true if CoverImage is a struct with all default values.
func (ci CoverImage) IsZero() bool {
	return ci == CoverImage{}
}

// URL returns the url of the cover image.
func (ci CoverImage) URL() string {
	return ci.url
}
