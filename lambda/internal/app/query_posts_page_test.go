package app

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueryPostsPageHandler_Handle(t *testing.T) {
	t.Parallel()
	customPageSize := 15

	testCases := map[string]struct {
		query         QueryPostsPage
		configureMock func(t *testing.T, m *MockPostService, q QueryPostsPage)

		shouldFail    bool
		expectedError string
	}{
		"returns first page when cursor is empty": {
			query: QueryPostsPage{PageSize: customPageSize},
			configureMock: func(_ *testing.T, m *MockPostService, q QueryPostsPage) {
				m.On("FirstPostsPage", mock.Anything, q.PageSize).Return(&PostsPage{}, nil)
			},
			shouldFail: false,
		},
		"uses cursor when provided": {
			query: QueryPostsPage{PageSize: customPageSize, Cursor: "test-cursor"},
			configureMock: func(_ *testing.T, m *MockPostService, q QueryPostsPage) {
				m.On("PostsPageFromCursor", mock.Anything, q.Cursor, q.PageSize).Return(&PostsPage{}, nil)
			},
			shouldFail: false,
		},
		"FirstPostsPage failed": {
			query: QueryPostsPage{PageSize: customPageSize},
			configureMock: func(_ *testing.T, m *MockPostService, q QueryPostsPage) {
				m.On("FirstPostsPage", mock.Anything, q.PageSize).Return(nil, errors.New("FirstPostsPage failed"))
			},
			shouldFail:    true,
			expectedError: "FirstPostsPage failed",
		},
		"PostsPageFromCursor failed": {
			query: QueryPostsPage{PageSize: customPageSize, Cursor: "test-cursor"},
			configureMock: func(_ *testing.T, m *MockPostService, q QueryPostsPage) {
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

			postService := NewMockPostService(t)
			tc.configureMock(t, postService, tc.query)

			handler := NewQueryPostsPageHandler(postService)
			_, err := handler.Handle(context.Background(), tc.query)

			if tc.shouldFail {
				assert.ErrorContains(t, err, tc.expectedError)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestQueryPostsPage_setMissingArguments(t *testing.T) {
	t.Parallel()

	t.Run("sets defaults", func(t *testing.T) {
		t.Parallel()

		query := QueryPostsPage{}
		expected := QueryPostsPage{PageSize: defaultPageSize}

		query.setMissingArguments()
		assert.Equal(t, expected, query)
	})

	t.Run("preserves values", func(t *testing.T) {
		t.Parallel()
		customPageSize := 15

		query := QueryPostsPage{PageSize: customPageSize}
		expected := QueryPostsPage{PageSize: customPageSize}

		query.setMissingArguments()
		assert.Equal(t, expected, query)
	})
}
