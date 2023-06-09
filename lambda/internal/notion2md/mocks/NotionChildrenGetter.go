// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	notionapi "github.com/jomei/notionapi"
)

// NotionChildrenGetter is an autogenerated mock type for the NotionChildrenGetter type
type NotionChildrenGetter struct {
	mock.Mock
}

// GetChildren provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotionChildrenGetter) GetChildren(_a0 context.Context, _a1 notionapi.BlockID, _a2 *notionapi.Pagination) (*notionapi.GetChildrenResponse, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *notionapi.GetChildrenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, notionapi.BlockID, *notionapi.Pagination) (*notionapi.GetChildrenResponse, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, notionapi.BlockID, *notionapi.Pagination) *notionapi.GetChildrenResponse); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notionapi.GetChildrenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, notionapi.BlockID, *notionapi.Pagination) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewNotionChildrenGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotionChildrenGetter creates a new instance of NotionChildrenGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotionChildrenGetter(t mockConstructorTestingTNewNotionChildrenGetter) *NotionChildrenGetter {
	mock := &NotionChildrenGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
