// Code generated by mockery v2.10.0. DO NOT EDIT.

package mock_gqlgen_example

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/wricardo/gqlgen_example/internal/gqlgen_example/model"
)

// UserResolver is an autogenerated mock type for the UserResolver type
type UserResolver struct {
	mock.Mock
}

type UserResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *UserResolver) EXPECT() *UserResolver_Expecter {
	return &UserResolver_Expecter{mock: &_m.Mock}
}

// CreatedBy provides a mock function with given fields: ctx, obj
func (_m *UserResolver) CreatedBy(ctx context.Context, obj *model.User) (*model.User, error) {
	ret := _m.Called(ctx, obj)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserResolver_CreatedBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatedBy'
type UserResolver_CreatedBy_Call struct {
	*mock.Call
}

// CreatedBy is a helper method to define mock.On call
//  - ctx context.Context
//  - obj *model.User
func (_e *UserResolver_Expecter) CreatedBy(ctx interface{}, obj interface{}) *UserResolver_CreatedBy_Call {
	return &UserResolver_CreatedBy_Call{Call: _e.mock.On("CreatedBy", ctx, obj)}
}

func (_c *UserResolver_CreatedBy_Call) Run(run func(ctx context.Context, obj *model.User)) *UserResolver_CreatedBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.User))
	})
	return _c
}

func (_c *UserResolver_CreatedBy_Call) Return(_a0 *model.User, _a1 error) *UserResolver_CreatedBy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Organization provides a mock function with given fields: ctx, obj
func (_m *UserResolver) Organization(ctx context.Context, obj *model.User) (*model.Organization, error) {
	ret := _m.Called(ctx, obj)

	var r0 *model.Organization
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.Organization); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserResolver_Organization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Organization'
type UserResolver_Organization_Call struct {
	*mock.Call
}

// Organization is a helper method to define mock.On call
//  - ctx context.Context
//  - obj *model.User
func (_e *UserResolver_Expecter) Organization(ctx interface{}, obj interface{}) *UserResolver_Organization_Call {
	return &UserResolver_Organization_Call{Call: _e.mock.On("Organization", ctx, obj)}
}

func (_c *UserResolver_Organization_Call) Run(run func(ctx context.Context, obj *model.User)) *UserResolver_Organization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.User))
	})
	return _c
}

func (_c *UserResolver_Organization_Call) Return(_a0 *model.Organization, _a1 error) *UserResolver_Organization_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdatedBy provides a mock function with given fields: ctx, obj
func (_m *UserResolver) UpdatedBy(ctx context.Context, obj *model.User) (*model.User, error) {
	ret := _m.Called(ctx, obj)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserResolver_UpdatedBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatedBy'
type UserResolver_UpdatedBy_Call struct {
	*mock.Call
}

// UpdatedBy is a helper method to define mock.On call
//  - ctx context.Context
//  - obj *model.User
func (_e *UserResolver_Expecter) UpdatedBy(ctx interface{}, obj interface{}) *UserResolver_UpdatedBy_Call {
	return &UserResolver_UpdatedBy_Call{Call: _e.mock.On("UpdatedBy", ctx, obj)}
}

func (_c *UserResolver_UpdatedBy_Call) Run(run func(ctx context.Context, obj *model.User)) *UserResolver_UpdatedBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.User))
	})
	return _c
}

func (_c *UserResolver_UpdatedBy_Call) Return(_a0 *model.User, _a1 error) *UserResolver_UpdatedBy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
