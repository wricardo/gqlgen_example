// Code generated by mockery v2.10.0. DO NOT EDIT.

package mock_gqlgen_example

import (
	mock "github.com/stretchr/testify/mock"
	generated "github.com/wricardo/gqlgen_example/internal/gqlgen_example/generated"
)

// ResolverRoot is an autogenerated mock type for the ResolverRoot type
type ResolverRoot struct {
	mock.Mock
}

type ResolverRoot_Expecter struct {
	mock *mock.Mock
}

func (_m *ResolverRoot) EXPECT() *ResolverRoot_Expecter {
	return &ResolverRoot_Expecter{mock: &_m.Mock}
}

// Mutation provides a mock function with given fields:
func (_m *ResolverRoot) Mutation() generated.MutationResolver {
	ret := _m.Called()

	var r0 generated.MutationResolver
	if rf, ok := ret.Get(0).(func() generated.MutationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.MutationResolver)
		}
	}

	return r0
}

// ResolverRoot_Mutation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mutation'
type ResolverRoot_Mutation_Call struct {
	*mock.Call
}

// Mutation is a helper method to define mock.On call
func (_e *ResolverRoot_Expecter) Mutation() *ResolverRoot_Mutation_Call {
	return &ResolverRoot_Mutation_Call{Call: _e.mock.On("Mutation")}
}

func (_c *ResolverRoot_Mutation_Call) Run(run func()) *ResolverRoot_Mutation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResolverRoot_Mutation_Call) Return(_a0 generated.MutationResolver) *ResolverRoot_Mutation_Call {
	_c.Call.Return(_a0)
	return _c
}

// Organization provides a mock function with given fields:
func (_m *ResolverRoot) Organization() generated.OrganizationResolver {
	ret := _m.Called()

	var r0 generated.OrganizationResolver
	if rf, ok := ret.Get(0).(func() generated.OrganizationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.OrganizationResolver)
		}
	}

	return r0
}

// ResolverRoot_Organization_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Organization'
type ResolverRoot_Organization_Call struct {
	*mock.Call
}

// Organization is a helper method to define mock.On call
func (_e *ResolverRoot_Expecter) Organization() *ResolverRoot_Organization_Call {
	return &ResolverRoot_Organization_Call{Call: _e.mock.On("Organization")}
}

func (_c *ResolverRoot_Organization_Call) Run(run func()) *ResolverRoot_Organization_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResolverRoot_Organization_Call) Return(_a0 generated.OrganizationResolver) *ResolverRoot_Organization_Call {
	_c.Call.Return(_a0)
	return _c
}

// Query provides a mock function with given fields:
func (_m *ResolverRoot) Query() generated.QueryResolver {
	ret := _m.Called()

	var r0 generated.QueryResolver
	if rf, ok := ret.Get(0).(func() generated.QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.QueryResolver)
		}
	}

	return r0
}

// ResolverRoot_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type ResolverRoot_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
func (_e *ResolverRoot_Expecter) Query() *ResolverRoot_Query_Call {
	return &ResolverRoot_Query_Call{Call: _e.mock.On("Query")}
}

func (_c *ResolverRoot_Query_Call) Run(run func()) *ResolverRoot_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResolverRoot_Query_Call) Return(_a0 generated.QueryResolver) *ResolverRoot_Query_Call {
	_c.Call.Return(_a0)
	return _c
}

// User provides a mock function with given fields:
func (_m *ResolverRoot) User() generated.UserResolver {
	ret := _m.Called()

	var r0 generated.UserResolver
	if rf, ok := ret.Get(0).(func() generated.UserResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.UserResolver)
		}
	}

	return r0
}

// ResolverRoot_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type ResolverRoot_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
func (_e *ResolverRoot_Expecter) User() *ResolverRoot_User_Call {
	return &ResolverRoot_User_Call{Call: _e.mock.On("User")}
}

func (_c *ResolverRoot_User_Call) Run(run func()) *ResolverRoot_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResolverRoot_User_Call) Return(_a0 generated.UserResolver) *ResolverRoot_User_Call {
	_c.Call.Return(_a0)
	return _c
}
