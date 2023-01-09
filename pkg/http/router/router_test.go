package router_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/framework/pkg/http/router"
	app_mock "github.com/supermetrolog/framework/tests/mocks/pkg/http/app"
	mock_handler "github.com/supermetrolog/framework/tests/mocks/pkg/http/interfaces/handler"
	router_mock "github.com/supermetrolog/framework/tests/mocks/pkg/http/router"
)

func TestGET(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPipelineMain := app_mock.NewMockPipeline(ctrl)
	mockPipelineLocal := app_mock.NewMockPipeline(ctrl)
	mockPipelineLocal.EXPECT().Pipe(mockPipelineMain)
	mockPipelineFactory := router_mock.NewMockPipelineFactory(ctrl)
	mockPipelineFactory.EXPECT().Create().Return(mockPipelineLocal)
	externalRouter := httprouter.New()
	// externalRouter.MethodNotAllowed.ServeHTTP()
	r := router.New(mockPipelineMain, mockPipelineFactory, externalRouter)

	mockHandler := mock_handler.NewMockHandler(ctrl)

	r.GET("/path", mockHandler)
}
