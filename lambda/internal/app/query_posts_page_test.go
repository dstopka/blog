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
	customPageSize := 15

	testCases := map[string]struct {
		query         app.QueryPostsPage
		configureMock func(t *testing.T, m *mocks.PostService, q app.QueryPostsPage)

		shouldFail    bool
		expectedError string
	}{
		"returns first page when cursor is empty": {
			query: app.QueryPostsPage{PageSize: customPageSize},
			configureMock: func(_ *testing.T, m *mocks.PostService, q app.QueryPostsPage) {
				m.On("FirstPostsPage", mock.Anything, q.PageSize).Return(&app.PostsPage{}, nil)
			},
			shouldFail: false,
		},
		"uses cursor when provided": {
			query: app.QueryPostsPage{PageSize: customPageSize, Cursor: "test-cursor"},
			configureMock: func(_ *testing.T, m *mocks.PostService, q app.QueryPostsPage) {
				m.On("PostsPageFromCursor", mock.Anything, q.Cursor, q.PageSize).Return(&app.PostsPage{}, nil)
			},
			shouldFail: false,
		},
		"FirstPostsPage failed": {
			query: app.QueryPostsPage{PageSize: customPageSize},
			configureMock: func(_ *testing.T, m *mocks.PostService, q app.QueryPostsPage) {
				m.On("FirstPostsPage", mock.Anything, q.PageSize).Return(nil, errors.New("FirstPostsPage failed"))
			},
			shouldFail:    true,
			expectedError: "FirstPostsPage failed",
		},
		"PostsPageFromCursor failed": {
			query: app.QueryPostsPage{PageSize: customPageSize, Cursor: "test-cursor"},
			configureMock: func(_ *testing.T, m *mocks.PostService, q app.QueryPostsPage) {
				m.On("PostsPageFromCursor", mock.Anything, q.Cursor, q.PageSize).Return(nil, errors.New("PostsPageFromCursor failed"))
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
			tc.configureMock(t, postService, tc.query)

			handler := app.NewQueryPostsPageHandler(postService)
			_, err := handler.Handle(context.Background(), tc.query)

			if tc.shouldFail {
				assert.ErrorContains(t, err, tc.expectedError)
				return
			}

			assert.NoError(t, err)
		})
	}
}