package goblack_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/supermetrolog/goblack"
	goblack_mock "github.com/supermetrolog/goblack/tests/mocks"
	mock_handler "github.com/supermetrolog/goblack/tests/mocks/pkg/http/interfaces/handler"
)

func TestPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRouter := goblack_mock.NewMockRouter(ctrl)
	mockMiddleware := mock_handler.NewMockMiddleware(ctrl)
	mockPipeline := goblack_mock.NewMockPipeline(ctrl)
	mockPipeline.EXPECT().Pipe(mockMiddleware)
	a := goblack.New(mockPipeline, mockRouter, goblack.ServerConfig{Addr: ":8080"})
	a.Pipe(mockMiddleware)
}

func TestGET(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMiddleware := mock_handler.NewMockMiddleware(ctrl)
	mockHandler := mock_handler.NewMockHandler(ctrl)
	mockPipeline := goblack_mock.NewMockPipeline(ctrl)
	mockRouter := goblack_mock.NewMockRouter(ctrl)
	mockRouter.EXPECT().GET("/path", mockHandler, mockMiddleware)
	a := goblack.New(mockPipeline, mockRouter, goblack.ServerConfig{Addr: ":8080"})
	a.GET("/path", mockHandler, mockMiddleware)
}
