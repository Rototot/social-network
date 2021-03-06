// Code generated by MockGen. DO NOT EDIT.
// Source: userRepository.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"
	users "social-network/pkg/users"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// FindByEmailAndPassword mocks base method.
func (m *MockUserRepositoryInterface) FindByEmailAndPassword(ctx context.Context, email string, password users.HashedPassword) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmailAndPassword", ctx, email, password)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmailAndPassword indicates an expected call of FindByEmailAndPassword.
func (mr *MockUserRepositoryInterfaceMockRecorder) FindByEmailAndPassword(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmailAndPassword", reflect.TypeOf((*MockUserRepositoryInterface)(nil).FindByEmailAndPassword), ctx, email, password)
}

// FindById mocks base method.
func (m *MockUserRepositoryInterface) FindById(ctx context.Context, id users.UserID) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserRepositoryInterfaceMockRecorder) FindById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).FindById), ctx, id)
}
