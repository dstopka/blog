package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dstopka/blog/lambda/internal/app"
	"github.com/dstopka/blog/lambda/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueryPostHandler_Handle(t *testing.T) {
	t.Parallel()
	query := app.QueryPost{"test-slug"}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		postService := mocks.NewPostService(t)
		postService.On("FindPostBySlug", mock.Anything, query.Slug).Return(&app.Post{}, nil)

		handler := app.NewQueryPostHandler(postService)
		_, err := handler.Handle(context.Background(), query)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		postService := mocks.NewPostService(t)
		postService.On("FindPostBySlug", mock.Anything, query.Slug).Return(nil, errors.New("FindPostBySlug failed"))

		handler := app.NewQueryPostHandler(postService)
		_, err := handler.Handle(context.Background(), query)

		assert.Error(t, err)
	})
}
