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

func TestQueryPostsPageHandler_Handle(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		configureMock func(t *testing.T, m *mocks.PostService)

		shouldFail    bool
		expectedError string
	}{
		"returns first page when cursor is empty": {
			configureMock: func(_ *testing.T, m *mocks.PostService) {
				m.On("AllPosts", mock.Anything).Return([]app.Post{}, nil)
			},
			shouldFail: false,
		},
		"uses cursor when provided": {
			configureMock: func(_ *testing.T, m *mocks.PostService) {
				m.On("AllPosts", mock.Anything).Return([]app.Post{}, nil)
			},
			shouldFail: false,
		},
		"FirstPostsPage failed": {
			configureMock: func(_ *testing.T, m *mocks.PostService) {
				m.On("AllPosts", mock.Anything).Return(nil, errors.New("FirstPostsPage failed"))
			},
			shouldFail:    true,
			expectedError: "FirstPostsPage failed",
		},
		"PostsPageFromCursor failed": {
			configureMock: func(_ *testing.T, m *mocks.PostService) {
				m.On("AllPosts", mock.Anything).Return(
					nil,
					errors.New("PostsPageFromCursor failed"),
				)
			},
			shouldFail:    true,
			expectedError: "PostsPageFromCursor failed",
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			postService := mocks.NewPostService(t)
			tc.configureMock(t, postService)

			handler := app.NewQueryAllPostsHandler(postService)
			_, err := handler.Handle(context.Background())

			if tc.shouldFail {
				assert.ErrorContains(t, err, tc.expectedError)
				return
			}

			assert.NoError(t, err)
		})
	}
}
