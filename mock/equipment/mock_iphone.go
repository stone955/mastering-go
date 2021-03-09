// Code generated by MockGen. DO NOT EDIT.
// Source: equipment/phone.go

// Package equipment is a generated GoMock package.
package equipment

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPhone is a mock of Phone interface.
type MockPhone struct {
	ctrl     *gomock.Controller
	recorder *MockPhoneMockRecorder
}

// MockPhoneMockRecorder is the mock recorder for MockPhone.
type MockPhoneMockRecorder struct {
	mock *MockPhone
}

// NewMockPhone creates a new mock instance.
func NewMockPhone(ctrl *gomock.Controller) *MockPhone {
	mock := &MockPhone{ctrl: ctrl}
	mock.recorder = &MockPhoneMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhone) EXPECT() *MockPhoneMockRecorder {
	return m.recorder
}

// DingTalk mocks base method.
func (m *MockPhone) DingTalk() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DingTalk")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DingTalk indicates an expected call of DingTalk.
func (mr *MockPhoneMockRecorder) DingTalk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DingTalk", reflect.TypeOf((*MockPhone)(nil).DingTalk))
}

// WeChat mocks base method.
func (m *MockPhone) WeChat() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeChat")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WeChat indicates an expected call of WeChat.
func (mr *MockPhoneMockRecorder) WeChat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeChat", reflect.TypeOf((*MockPhone)(nil).WeChat))
}

// ZhiHu mocks base method.
func (m *MockPhone) ZhiHu() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZhiHu")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ZhiHu indicates an expected call of ZhiHu.
func (mr *MockPhoneMockRecorder) ZhiHu() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZhiHu", reflect.TypeOf((*MockPhone)(nil).ZhiHu))
}