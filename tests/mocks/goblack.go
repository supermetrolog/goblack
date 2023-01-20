// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/../goblack.go

// Package mock_goblack is a generated GoMock package.
package mock_goblack

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	handler "github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	httpcontext "github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
)

// MockPipeline is a mock of Pipeline interface.
type MockPipeline struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineMockRecorder
}

// MockPipelineMockRecorder is the mock recorder for MockPipeline.
type MockPipelineMockRecorder struct {
	mock *MockPipeline
}

// NewMockPipeline creates a new mock instance.
func NewMockPipeline(ctrl *gomock.Controller) *MockPipeline {
	mock := &MockPipeline{ctrl: ctrl}
	mock.recorder = &MockPipelineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipeline) EXPECT() *MockPipelineMockRecorder {
	return m.recorder
}

// Handler mocks base method.
func (m *MockPipeline) Handler(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", c, next)
	ret0, _ := ret[0].(httpcontext.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handler indicates an expected call of Handler.
func (mr *MockPipelineMockRecorder) Handler(c, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockPipeline)(nil).Handler), c, next)
}

// Pipe mocks base method.
func (m *MockPipeline) Pipe(arg0 handler.Middleware) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Pipe", arg0)
}

// Pipe indicates an expected call of Pipe.
func (mr *MockPipelineMockRecorder) Pipe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockPipeline)(nil).Pipe), arg0)
}

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// GET mocks base method.
func (m *MockRouter) GET(path string, handler handler.Handler, middlewares ...handler.Middleware) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path, handler}
	for _, a := range middlewares {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "GET", varargs...)
}

// GET indicates an expected call of GET.
func (mr *MockRouterMockRecorder) GET(path, handler interface{}, middlewares ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, handler}, middlewares...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockRouter)(nil).GET), varargs...)
}

// ServeHTTP mocks base method.
func (m *MockRouter) ServeHTTP(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", arg0, arg1)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockRouterMockRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockRouter)(nil).ServeHTTP), arg0, arg1)
}