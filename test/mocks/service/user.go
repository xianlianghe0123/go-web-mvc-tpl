// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	do "github.com/xianlianghe0123/go-web-mvc-tpl/internal/model/do"
	dto "github.com/xianlianghe0123/go-web-mvc-tpl/internal/model/dto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, req)
}

// Register mocks base method.
func (m *MockUserService) Register(ctx context.Context, req *dto.RegisterRequest) (*do.UserDO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, req)
	ret0, _ := ret[0].(*do.UserDO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceMockRecorder) Register(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), ctx, req)
}
