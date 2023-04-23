// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/dstopka/blog/lambda/internal/app"

	mock "github.com/stretchr/testify/mock"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// FindPostBySlug provides a mock function with given fields: ctx, slug
func (_m *PostService) FindPostBySlug(ctx context.Context, slug string) (*app.Post, error) {
	ret := _m.Called(ctx, slug)

	var r0 *app.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*app.Post, error)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *app.Post); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app.Post)
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
func (_m *PostService) FirstPostsPage(ctx context.Context, pageSize int) (*app.PostsPage, error) {
	ret := _m.Called(ctx, pageSize)

	var r0 *app.PostsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*app.PostsPage, error)); ok {
		return rf(ctx, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *app.PostsPage); ok {
		r0 = rf(ctx, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app.PostsPage)
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
func (_m *PostService) PostsPageFromCursor(ctx context.Context, cursor string, pageSize int) (*app.PostsPage, error) {
	ret := _m.Called(ctx, cursor, pageSize)

	var r0 *app.PostsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (*app.PostsPage, error)); ok {
		return rf(ctx, cursor, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *app.PostsPage); ok {
		r0 = rf(ctx, cursor, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app.PostsPage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, cursor, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostService(t mockConstructorTestingTNewPostService) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}