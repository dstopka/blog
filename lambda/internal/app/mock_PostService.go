// Code generated by mockery v2.24.0. DO NOT EDIT.

package app

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockPostService is an autogenerated mock type for the PostService type
type MockPostService struct {
	mock.Mock
}

// FindPostBySlug provides a mock function with given fields: ctx, slug
func (_m *MockPostService) FindPostBySlug(ctx context.Context, slug string) (*Post, error) {
	ret := _m.Called(ctx, slug)

	var r0 *Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Post, error)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Post); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPostsPage provides a mock function with given fields: ctx, pageSize
func (_m *MockPostService) FirstPostsPage(ctx context.Context, pageSize int) (*PostsPage, error) {
	ret := _m.Called(ctx, pageSize)

	var r0 *PostsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*PostsPage, error)); ok {
		return rf(ctx, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *PostsPage); ok {
		r0 = rf(ctx, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PostsPage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostsPageFromCursor provides a mock function with given fields: ctx, cursor, pageSize
func (_m *MockPostService) PostsPageFromCursor(ctx context.Context, cursor string, pageSize int) (*PostsPage, error) {
	ret := _m.Called(ctx, cursor, pageSize)

	var r0 *PostsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (*PostsPage, error)); ok {
		return rf(ctx, cursor, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *PostsPage); ok {
		r0 = rf(ctx, cursor, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PostsPage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, cursor, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockPostService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPostService creates a new instance of MockPostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPostService(t mockConstructorTestingTNewMockPostService) *MockPostService {
	mock := &MockPostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}