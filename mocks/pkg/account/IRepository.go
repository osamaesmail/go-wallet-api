// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	account "go-api-grpc/pkg/account"

	mock "github.com/stretchr/testify/mock"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: model
func (_m *IRepository) Create(model account.Account) (account.Account, error) {
	ret := _m.Called(model)

	var r0 account.Account
	if rf, ok := ret.Get(0).(func(account.Account) account.Account); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Get(0).(account.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(account.Account) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: req
func (_m *IRepository) List(req account.ListRequest) ([]account.Account, error) {
	ret := _m.Called(req)

	var r0 []account.Account
	if rf, ok := ret.Get(0).(func(account.ListRequest) []account.Account); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(account.ListRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepository(t mockConstructorTestingTNewIRepository) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
