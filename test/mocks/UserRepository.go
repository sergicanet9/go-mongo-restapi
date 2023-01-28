// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, entity
func (_m *UserRepository) Create(ctx context.Context, entity interface{}) (string, error) {
	ret := _m.Called(ctx, entity)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) string); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMany provides a mock function with given fields: ctx, entities
func (_m *UserRepository) CreateMany(ctx context.Context, entities []interface{}) ([]string, error) {
	ret := _m.Called(ctx, entities)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) []string); ok {
		r0 = rf(ctx, entities)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []interface{}) error); ok {
		r1 = rf(ctx, entities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, ID
func (_m *UserRepository) Delete(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, filter, skip, take
func (_m *UserRepository) Get(ctx context.Context, filter map[string]interface{}, skip *int, take *int) ([]interface{}, error) {
	ret := _m.Called(ctx, filter, skip, take)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, *int, *int) []interface{}); ok {
		r0 = rf(ctx, filter, skip, take)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}, *int, *int) error); ok {
		r1 = rf(ctx, filter, skip, take)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *UserRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	ret := _m.Called(ctx, ID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string) interface{}); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ID, entity
func (_m *UserRepository) Update(ctx context.Context, ID string, entity interface{}) error {
	ret := _m.Called(ctx, ID, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, ID, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
