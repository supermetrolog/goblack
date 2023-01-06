package pipeline_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
	"github.com/supermetrolog/framework/pkg/http/pipeline"
	mock_request "github.com/supermetrolog/framework/tests/mocks/pkg/http/interfaces/request"
	mock_response "github.com/supermetrolog/framework/tests/mocks/pkg/http/interfaces/response"
	mock_pipeline "github.com/supermetrolog/framework/tests/mocks/pkg/http/pipeline"
)

func TestPipeline_pipe(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mock_pipeline.NewMockMiddleware(ctrl)
	mockHandler2 := mock_pipeline.NewMockMiddleware(ctrl)

	p.Pipe(mockHandler)
	p.Pipe(mockHandler2)

	assert.NotEmpty(t, p.Handlers)
	assert.Equal(t, 2, p.Handlers.Length())
}

func TestPipeline_runWithDefaultHandler(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReq := mock_request.NewMockRequest(ctrl)
	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockRes := mock_response.NewMockResponse(ctrl)
	mockHandler := mock_pipeline.NewMockHandler(ctrl)
	mockHandler.EXPECT().Handler(mockResWriter, mockReq).Return(mockRes, nil)

	_, err := p.Handler(mockResWriter, mockReq, mockHandler)

	assert.NoError(t, err)
}

func TestPipeline_runWithNilDefaultHandler(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := mock_request.NewMockRequest(ctrl)
	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	_, err := p.Handler(mockResWriter, mockReq, nil)
	assert.Error(t, err)
}

func TestPipeline_runWithManyHandlers(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	firstCall := mockResWriter.EXPECT().AddHeader("header1", "value1")
	secondCall := mockResWriter.EXPECT().AddHeader("header2", "value2")
	thirdCall := mockResWriter.EXPECT().AddHeader("header4", "value4")
	gomock.InOrder(
		firstCall,
		secondCall,
		thirdCall,
	)

	mockResWriter.EXPECT().SetContent("content")
	mockResWriter.EXPECT().Response()

	mockReq := mock_request.NewMockRequest(ctrl)

	mock1 := mockMiddleware1{}
	mock2 := mockMiddleware2{}
	last := mockHandler{}

	p.Pipe(mock1)
	p.Pipe(mock2)
	_, err := p.Handler(mockResWriter, mockReq, last)
	assert.NoError(t, err)
}

func TestPipeline_doubleRun(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockReq := mock_request.NewMockRequest(ctrl)

	mockHandlerDefault := mock_pipeline.NewMockHandler(ctrl)
	mockHandlerDefault.EXPECT().Handler(mockResWriter, mockReq).Times(2)

	mockHandler2 := mock_pipeline.NewMockMiddleware(ctrl)
	mockHandler2.EXPECT().Handler(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
		return next.Handler(res, req)
	}).Times(2)

	mockHandler1 := mock_pipeline.NewMockMiddleware(ctrl)
	mockHandler1.EXPECT().Handler(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
		return next.Handler(res, req)
	}).Times(2)

	p.Pipe(mockHandler1)
	p.Pipe(mockHandler2)

	_, err := p.Handler(mockResWriter, mockReq, mockHandlerDefault)
	assert.NoError(t, err)

	_, err = p.Handler(mockResWriter, mockReq, mockHandlerDefault)
	assert.NoError(t, err)
}

func TestPipeline_PipelineInPipeline(t *testing.T) {
	p1 := pipeline.New()
	p2 := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockReq := mock_request.NewMockRequest(ctrl)

	mockHandlerDefault := mock_pipeline.NewMockHandler(ctrl)
	mockHandlerDefault.EXPECT().Handler(mockResWriter, mockReq)

	mockHandler2 := mock_pipeline.NewMockMiddleware(ctrl)
	mockHandler2.EXPECT().Handler(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
		return next.Handler(res, req)
	}).Times(2)

	mockHandler1 := mock_pipeline.NewMockMiddleware(ctrl)
	mockHandler1.EXPECT().Handler(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
		return next.Handler(res, req)
	}).Times(2)

	p1.Pipe(mockHandler1)
	p1.Pipe(mockHandler2)

	p2.Pipe(mockHandler1)
	p2.Pipe(mockHandler2)

	p1.Pipe(p2)
	_, err := p1.Handler(mockResWriter, mockReq, mockHandlerDefault)
	assert.NoError(t, err)
}

type mockMiddleware1 struct{}

func (m mockMiddleware1) Handler(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
	res.AddHeader("header1", "value1")
	return next.Handler(res, req)
}

type mockMiddleware2 struct{}

func (m mockMiddleware2) Handler(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
	res.AddHeader("header2", "value2")
	return next.Handler(res, req)
}

type mockMiddleware3 struct{}

func (m mockMiddleware3) Handler(res response.ResponseWriter, req request.Request, next pipeline.Handler) (response.Response, error) {
	res.AddHeader("header3", "value3")
	res.SetContent("suka")
	return res.Response(), nil
}

type mockHandler struct{}

func (m mockHandler) Handler(res response.ResponseWriter, req request.Request) (response.Response, error) {
	res.AddHeader("header4", "value4")
	res.SetContent("content")
	return res.Response(), nil
}
