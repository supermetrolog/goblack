package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/supermetrolog/framework/pkg/http/app"
	app_mock "github.com/supermetrolog/framework/tests/mocks/pkg/http/app"
	mock_handler "github.com/supermetrolog/framework/tests/mocks/pkg/http/interfaces/handler"
)

func TestPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRouter := app_mock.NewMockRouter(ctrl)
	mockMiddleware := mock_handler.NewMockMiddleware(ctrl)
	mockPipeline := app_mock.NewMockPipeline(ctrl)
	mockPipeline.EXPECT().Pipe(mockMiddleware)
	a := app.New(mockPipeline, mockRouter)
	a.Pipe(mockMiddleware)
}

func TestGET(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMiddleware := mock_handler.NewMockMiddleware(ctrl)
	mockHandler := mock_handler.NewMockHandler(ctrl)
	mockPipeline := app_mock.NewMockPipeline(ctrl)
	mockRouter := app_mock.NewMockRouter(ctrl)
	mockRouter.EXPECT().GET("/path", mockHandler, mockMiddleware)
	a := app.New(mockPipeline, mockRouter)
	a.GET("/path", mockHandler, mockMiddleware)
}
