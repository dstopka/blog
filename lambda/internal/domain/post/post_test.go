package post_test

import (
	"testing"
	"time"

	"github.com/dstopka/blog/lambda/internal/domain/post"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	postUUID := uuid.NewString()
	title := "Test post"
	slug := "test-post"
	description := "test description"
	publishedTime := time.Now()
	coverImg, err := post.NewCoverImage("http://s3.amazonaws.com/bucket/test-post-cover")
	require.NoError(t, err)

	post, err := post.New(
		postUUID,
		title,
		slug,
		description,
		publishedTime,
		coverImg,
	)
	require.NoError(t, err)

	assert.Equal(t, postUUID, post.ID())
	assert.Equal(t, title, post.Title())
	assert.Equal(t, slug, post.Slug())
	assert.Equal(t, description, post.Description())
	assert.Equal(t, coverImg, post.CoverImage())
	assert.Equal(t, publishedTime, post.PublishedTime())
}

func TestNew_invalid(t *testing.T) {
	t.Parallel()

	postUUID := uuid.NewString()
	title := "Test post"
	slug := "test-post"
	description := "Test post's description"
	publishedTime := time.Now()
	coverImg := createExampleCoverImage(t)

	testCases := map[string]struct {
		createPost func() (*post.Post, error)
		withError  error
	}{
		"empty uuid": {
			createPost: func() (*post.Post, error) {
				return post.New("", title, slug, description, publishedTime, coverImg)
			},
			withError: post.ErrEmptyUUID,
		},
		"empty title": {
			createPost: func() (*post.Post, error) {
				return post.New(postUUID, "", slug, description, publishedTime, coverImg)
			},
			withError: post.ErrEmptyTitle,
		},
		"empty slug": {
			createPost: func() (*post.Post, error) {
				return post.New(postUUID, title, "", description, publishedTime, coverImg)
			},
			withError: post.ErrEmptySlug,
		},
		"empty description": {
			createPost: func() (*post.Post, error) {
				return post.New(postUUID, title, slug, "", publishedTime, coverImg)
			},
			withError: post.ErrEmptyDescription,
		},
		"zero published time": {
			createPost: func() (*post.Post, error) {
				return post.New(postUUID, title, slug, description, time.Time{}, coverImg)
			},
			withError: post.ErrZeroPublishedTime,
		},
		"zero cover image": {
			createPost: func() (*post.Post, error) {
				return post.New(postUUID, title, slug, description, publishedTime, post.CoverImage{})
			},
			withError: post.ErrZeroCoverImage,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := tc.createPost()
			assert.ErrorIs(t, err, tc.withError)
		})
	}
}
