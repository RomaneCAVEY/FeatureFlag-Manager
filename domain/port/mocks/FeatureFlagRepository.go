// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/port/featureFlagManagerPort.go

// Package mock_port is a generated GoMock package.
package mock_port

import (
	reflect "reflect"

	dto "github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	entities "github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockFeatureFlagManagerPort is a mock of FeatureFlagManagerPort interface.
type MockFeatureFlagManagerPort struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureFlagManagerPortMockRecorder
}

// MockFeatureFlagManagerPortMockRecorder is the mock recorder for MockFeatureFlagManagerPort.
type MockFeatureFlagManagerPortMockRecorder struct {
	mock *MockFeatureFlagManagerPort
}

// NewMockFeatureFlagManagerPort creates a new mock instance.
func NewMockFeatureFlagManagerPort(ctrl *gomock.Controller) *MockFeatureFlagManagerPort {
	mock := &MockFeatureFlagManagerPort{ctrl: ctrl}
	mock.recorder = &MockFeatureFlagManagerPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureFlagManagerPort) EXPECT() *MockFeatureFlagManagerPortMockRecorder {
	return m.recorder
}

// CreateAFeatureFlag mocks base method.
func (m *MockFeatureFlagManagerPort) CreateAFeatureFlag(arg0 dto.CreateAFeatureFlagDTO, arg1 entities.User) (*entities.FeatureFlag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAFeatureFlag", arg0, arg1)
	ret0, _ := ret[0].(*entities.FeatureFlag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAFeatureFlag indicates an expected call of CreateAFeatureFlag.
func (mr *MockFeatureFlagManagerPortMockRecorder) CreateAFeatureFlag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAFeatureFlag", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).CreateAFeatureFlag), arg0, arg1)
}

// DeleteFeatureFlag mocks base method.
func (m *MockFeatureFlagManagerPort) DeleteFeatureFlag(arg0 uint32) (*entities.FeatureFlag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFeatureFlag", arg0)
	ret0, _ := ret[0].(*entities.FeatureFlag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFeatureFlag indicates an expected call of DeleteFeatureFlag.
func (mr *MockFeatureFlagManagerPortMockRecorder) DeleteFeatureFlag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFeatureFlag", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).DeleteFeatureFlag), arg0)
}

// GetAllFeatureFlags mocks base method.
func (m *MockFeatureFlagManagerPort) GetAllFeatureFlags(arg0, arg1 int) (*[]entities.FeatureFlag, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeatureFlags", arg0, arg1)
	ret0, _ := ret[0].(*[]entities.FeatureFlag)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllFeatureFlags indicates an expected call of GetAllFeatureFlags.
func (mr *MockFeatureFlagManagerPortMockRecorder) GetAllFeatureFlags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeatureFlags", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).GetAllFeatureFlags), arg0, arg1)
}

// GetFeatureFlagsByApplication mocks base method.
func (m *MockFeatureFlagManagerPort) GetFeatureFlagsByApplication(arg0 string, arg1, arg2 int) (*[]entities.FeatureFlag, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureFlagsByApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]entities.FeatureFlag)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFeatureFlagsByApplication indicates an expected call of GetFeatureFlagsByApplication.
func (mr *MockFeatureFlagManagerPortMockRecorder) GetFeatureFlagsByApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureFlagsByApplication", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).GetFeatureFlagsByApplication), arg0, arg1, arg2)
}

// GetFeatureFlagsById mocks base method.
func (m *MockFeatureFlagManagerPort) GetFeatureFlagsById(arg0 uint32) (*entities.FeatureFlag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureFlagsById", arg0)
	ret0, _ := ret[0].(*entities.FeatureFlag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureFlagsById indicates an expected call of GetFeatureFlagsById.
func (mr *MockFeatureFlagManagerPortMockRecorder) GetFeatureFlagsById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureFlagsById", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).GetFeatureFlagsById), arg0)
}

// ModifyFeatureFlag mocks base method.
func (m *MockFeatureFlagManagerPort) ModifyFeatureFlag(arg0 uint32, arg1 dto.ModifyFeatureFlagDTO, arg2 entities.User) (*entities.FeatureFlag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyFeatureFlag", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entities.FeatureFlag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyFeatureFlag indicates an expected call of ModifyFeatureFlag.
func (mr *MockFeatureFlagManagerPortMockRecorder) ModifyFeatureFlag(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyFeatureFlag", reflect.TypeOf((*MockFeatureFlagManagerPort)(nil).ModifyFeatureFlag), arg0, arg1, arg2)
}
