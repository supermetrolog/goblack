// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/http/interfaces/request/request.go

// Package mock_request is a generated GoMock package.
package mock_request

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockRequest) Body() io.ReadCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].(io.ReadCloser)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockRequest)(nil).Body))
}

// Header mocks base method.
func (m *MockRequest) Header(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockRequestMockRecorder) Header(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockRequest)(nil).Header), name)
}

// Headers mocks base method.
func (m *MockRequest) Headers() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Headers")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Headers indicates an expected call of Headers.
func (mr *MockRequestMockRecorder) Headers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Headers", reflect.TypeOf((*MockRequest)(nil).Headers))
}

// QueryParam mocks base method.
func (m *MockRequest) QueryParam(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParam", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryParam indicates an expected call of QueryParam.
func (mr *MockRequestMockRecorder) QueryParam(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParam", reflect.TypeOf((*MockRequest)(nil).QueryParam), name)
}

// QueryParams mocks base method.
func (m *MockRequest) QueryParams() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParams")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// QueryParams indicates an expected call of QueryParams.
func (mr *MockRequestMockRecorder) QueryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParams", reflect.TypeOf((*MockRequest)(nil).QueryParams))
}
