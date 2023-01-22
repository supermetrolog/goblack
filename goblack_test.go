package goblack_test

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/supermetrolog/goblack"
	mock_goblack "github.com/supermetrolog/goblack/mocks"
)

func TestPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRouter := mock_goblack.NewMockRouter(ctrl)
	mockMiddleware := mock_goblack.NewMockMiddleware(ctrl)
	mockPipeline := mock_goblack.NewMockPipeline(ctrl)
	mockPipeline.EXPECT().Pipe(mockMiddleware)
	a := goblack.New(mockPipeline, mockRouter, goblack.ServerConfig{Addr: ":8080"})
	a.Pipe(mockMiddleware)
}

func TestRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMiddleware := mock_goblack.NewMockMiddleware(ctrl)
	mockHandler := mock_goblack.NewMockHandler(ctrl)
	mockPipeline := mock_goblack.NewMockPipeline(ctrl)
	mockRouter := mock_goblack.NewMockRouter(ctrl)

	mockRouter.EXPECT().GET("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().POST("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().PUT("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().PATCH("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().DELETE("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().OPTIONS("/path", mockHandler, mockMiddleware)
	mockRouter.EXPECT().HEAD("/path", mockHandler, mockMiddleware)

	testRequest := httptest.NewRequest("POST", "/path", nil)
	mockRouter.EXPECT().ServeHTTP(nil, testRequest)

	a := goblack.New(mockPipeline, mockRouter, goblack.ServerConfig{Addr: ":8080"})

	a.GET("/path", mockHandler, mockMiddleware)
	a.POST("/path", mockHandler, mockMiddleware)
	a.PUT("/path", mockHandler, mockMiddleware)
	a.PATCH("/path", mockHandler, mockMiddleware)
	a.DELETE("/path", mockHandler, mockMiddleware)
	a.OPTIONS("/path", mockHandler, mockMiddleware)
	a.HEAD("/path", mockHandler, mockMiddleware)
	a.ServeHTTP(nil, testRequest)
}

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRouter := mock_goblack.NewMockRouter(ctrl)
	mockHttpCondext := mock_goblack.NewMockContext(ctrl)
	mockPipeline := mock_goblack.NewMockPipeline(ctrl)
	mockNextHandler := mock_goblack.NewMockHandler(ctrl)
	mockPipeline.EXPECT().Handler(mockHttpCondext, mockNextHandler)
	a := goblack.New(mockPipeline, mockRouter, goblack.ServerConfig{Addr: ":8080"})
	a.Handler(mockHttpCondext, mockNextHandler)
}
