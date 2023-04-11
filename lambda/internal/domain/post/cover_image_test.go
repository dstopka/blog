package post_test

import (
	"testing"

	"github.com/dstopka/blog/lambda/internal/domain/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCoverImage(t *testing.T) {
	t.Parallel()

	url := "http://s3.amazonaws.com/bucket/test-cover-image"

	coverImg, err := post.NewCoverImage(url)
	require.NoError(t, err)

	assert.Equal(t, url, coverImg.URL())

	_, err = post.NewCoverImage("")
	assert.ErrorIs(t, err, post.ErrEmptyCoverImageURL)
}

func TestCoverImage_IsZero(t *testing.T) {
	t.Parallel()

	coverImg := createExampleCoverImage(t)
	assert.False(t, coverImg.IsZero())

	coverImg = post.CoverImage{}
	assert.True(t, coverImg.IsZero())
}

func createExampleCoverImage(t *testing.T) post.CoverImage {
	cover, err := post.NewCoverImage("http://s3.amazonaws.com/bucket/example-post-cover")
	require.NoError(t, err)

	return cover
}
