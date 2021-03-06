// Code generated by mockery v2.10.0. DO NOT EDIT.

package mock_gqlgen_example

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/wricardo/gqlgen_example/internal/gqlgen_example/model"
)

// QueryResolver is an autogenerated mock type for the QueryResolver type
type QueryResolver struct {
	mock.Mock
}

type QueryResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryResolver) EXPECT() *QueryResolver_Expecter {
	return &QueryResolver_Expecter{mock: &_m.Mock}
}

// Healthcheck provides a mock function with given fields: ctx
func (_m *QueryResolver) Healthcheck(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryResolver_Healthcheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Healthcheck'
type QueryResolver_Healthcheck_Call struct {
	*mock.Call
}

// Healthcheck is a helper method to define mock.On call
//  - ctx context.Context
func (_e *QueryResolver_Expecter) Healthcheck(ctx interface{}) *QueryResolver_Healthcheck_Call {
	return &QueryResolver_Healthcheck_Call{Call: _e.mock.On("Healthcheck", ctx)}
}

func (_c *QueryResolver_Healthcheck_Call) Run(run func(ctx context.Context)) *QueryResolver_Healthcheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QueryResolver_Healthcheck_Call) Return(_a0 string, _a1 error) *QueryResolver_Healthcheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Users provides a mock function with given fields: ctx, page, size
func (_m *QueryResolver) Users(ctx context.Context, page *int, size *int) (*model.UsersPage, error) {
	ret := _m.Called(ctx, page, size)

	var r0 *model.UsersPage
	if rf, ok := ret.Get(0).(func(context.Context, *int, *int) *model.UsersPage); ok {
		r0 = rf(ctx, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UsersPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int, *int) error); ok {
		r1 = rf(ctx, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryResolver_Users_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Users'
type QueryResolver_Users_Call struct {
	*mock.Call
}

// Users is a helper method to define mock.On call
//  - ctx context.Context
//  - page *int
//  - size *int
func (_e *QueryResolver_Expecter) Users(ctx interface{}, page interface{}, size interface{}) *QueryResolver_Users_Call {
	return &QueryResolver_Users_Call{Call: _e.mock.On("Users", ctx, page, size)}
}

func (_c *QueryResolver_Users_Call) Run(run func(ctx context.Context, page *int, size *int)) *QueryResolver_Users_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*int), args[2].(*int))
	})
	return _c
}

func (_c *QueryResolver_Users_Call) Return(_a0 *model.UsersPage, _a1 error) *QueryResolver_Users_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
