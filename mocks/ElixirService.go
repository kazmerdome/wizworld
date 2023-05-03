// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	context "context"

	elixir "github.com/kazmerdome/wizworld/internal/elixir"
	mock "github.com/stretchr/testify/mock"
)

// ElixirService is an autogenerated mock type for the ElixirService type
type ElixirService struct {
	mock.Mock
}

// GetElixirsByIngredients provides a mock function with given fields: ctx, ingredients
func (_m *ElixirService) GetElixirsByIngredients(ctx context.Context, ingredients []string) ([]elixir.Elixir, error) {
	ret := _m.Called(ctx, ingredients)

	var r0 []elixir.Elixir
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]elixir.Elixir, error)); ok {
		return rf(ctx, ingredients)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []elixir.Elixir); ok {
		r0 = rf(ctx, ingredients)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]elixir.Elixir)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ingredients)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewElixirService interface {
	mock.TestingT
	Cleanup(func())
}

// NewElixirService creates a new instance of ElixirService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewElixirService(t mockConstructorTestingTNewElixirService) *ElixirService {
	mock := &ElixirService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
