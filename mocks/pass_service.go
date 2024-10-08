// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	jwt "github.com/dgrijalva/jwt-go"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// PasswordService is an autogenerated mock type for the PasswordService type
type PasswordService struct {
	mock.Mock
}

// PasswordComparator provides a mock function with given fields: hash, password
func (_m *PasswordService) PasswordComparator(hash string, password string) bool {
	ret := _m.Called(hash, password)

	if len(ret) == 0 {
		panic("no return value specified for PasswordComparator")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hash, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PasswordHasher provides a mock function with given fields: password
func (_m *PasswordService) PasswordHasher(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for PasswordHasher")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenClaimer provides a mock function with given fields: tokenstr
func (_m *PasswordService) TokenClaimer(tokenstr string) (*jwt.Token, error) {
	ret := _m.Called(tokenstr)

	if len(ret) == 0 {
		panic("no return value specified for TokenClaimer")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(tokenstr)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(tokenstr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenstr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenGenerator provides a mock function with given fields: id, email, isadmin
func (_m *PasswordService) TokenGenerator(id primitive.ObjectID, email string, isadmin bool) (string, error) {
	ret := _m.Called(id, email, isadmin)

	if len(ret) == 0 {
		panic("no return value specified for TokenGenerator")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, string, bool) (string, error)); ok {
		return rf(id, email, isadmin)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, string, bool) string); ok {
		r0 = rf(id, email, isadmin)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID, string, bool) error); ok {
		r1 = rf(id, email, isadmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPasswordService creates a new instance of PasswordService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordService {
	mock := &PasswordService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
