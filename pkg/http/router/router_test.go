package router_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	goblack_mock "github.com/supermetrolog/goblack/mocks"
	"github.com/supermetrolog/goblack/pkg/http/router"
	mock_router "github.com/supermetrolog/goblack/pkg/http/router/mocks"
)

func TestMethods(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPipelineMain := goblack_mock.NewMockPipeline(ctrl)
	mockPipelineLocal := goblack_mock.NewMockPipeline(ctrl)
	mockPipelineLocal.EXPECT().Pipe(mockPipelineMain).Times(7)
	mockPipelineFactory := mock_router.NewMockPipelineFactory(ctrl)
	mockPipelineFactory.EXPECT().Create().Return(mockPipelineLocal).Times(7)
	externalRouter := httprouter.New()
	r := router.New(mockPipelineMain, mockPipelineFactory, externalRouter)

	mockHandler := goblack_mock.NewMockHandler(ctrl)

	r.GET("/path", mockHandler)
	r.POST("/path", mockHandler)
	r.PUT("/path", mockHandler)
	r.PATCH("/path", mockHandler)
	r.DELETE("/path", mockHandler)
	r.OPTIONS("/path", mockHandler)
	r.HEAD("/path", mockHandler)
}
