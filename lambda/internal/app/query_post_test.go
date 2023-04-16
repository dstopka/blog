package app

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueryPostHandler_Handle(t *testing.T) {
	t.Parallel()
	query := QueryPost{"test-slug"}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		postService := NewMockPostService(t)
		postService.On("FindPostBySlug", mock.Anything, query.Slug).Return(&Post{}, nil)

		handler := NewQueryPostHandler(postService)
		_, err := handler.Handle(context.Background(), query)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		postService := NewMockPostService(t)
		postService.On("FindPostBySlug", mock.Anything, query.Slug).Return(nil, errors.New("FindPostBySlug failed"))

		handler := NewQueryPostHandler(postService)
		_, err := handler.Handle(context.Background(), query)

		assert.Error(t, err)
	})
}
