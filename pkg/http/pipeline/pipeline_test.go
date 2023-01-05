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

	mockHandle := mock_pipeline.NewMockHandle(ctrl)
	mockHandle2 := mock_pipeline.NewMockHandle(ctrl)

	p.Pipe(mockHandle)
	p.Pipe(mockHandle2)

	assert.NotEmpty(t, p.Handlers)
	assert.Equal(t, 2, p.Handlers.Length())
}

func TestPipeline_runWithDefaultHandle(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReq := mock_request.NewMockRequest(ctrl)
	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockRes := mock_response.NewMockResponse(ctrl)
	mockHandle := mock_pipeline.NewMockHandle(ctrl)
	mockHandle.EXPECT().Handle(mockResWriter, mockReq, nil).Return(mockRes, nil)

	_, err := p.Handle(mockResWriter, mockReq, mockHandle)

	assert.NoError(t, err)
}

func TestPipeline_runWithNilDefaultHandle(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := mock_request.NewMockRequest(ctrl)
	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	_, err := p.Handle(mockResWriter, mockReq, nil)
	assert.Error(t, err)
}

func TestPipeline_runWithManyHandlers(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	firstCall := mockResWriter.EXPECT().AddHeader(map[string]string{"header1": "value1"})
	secondCall := mockResWriter.EXPECT().AddHeader(map[string]string{"header2": "value2"})
	thirdCall := mockResWriter.EXPECT().AddHeader(map[string]string{"header3": "value3"})
	gomock.InOrder(
		firstCall,
		secondCall,
		thirdCall,
	)
	mockResWriter.EXPECT().SetContent("suka")
	mockResWriter.EXPECT().Response()

	mockReq := mock_request.NewMockRequest(ctrl)

	mock1 := mockMiddleware1{}
	mock2 := mockMiddleware2{}
	last := mockMiddleware3{}

	p.Pipe(mock1)
	p.Pipe(mock2)
	_, err := p.Handle(mockResWriter, mockReq, last)
	assert.NoError(t, err)
}

func TestPipeline_doubleRun(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockReq := mock_request.NewMockRequest(ctrl)

	mockHandleDefault := mock_pipeline.NewMockHandle(ctrl)
	mockHandleDefault.EXPECT().Handle(mockResWriter, mockReq, nil).Times(2)

	mockHandle2 := mock_pipeline.NewMockHandle(ctrl)
	mockHandle2.EXPECT().Handle(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
		return next.Handle(res, req, nil)
	}).Times(2)

	mockHandle1 := mock_pipeline.NewMockHandle(ctrl)
	mockHandle1.EXPECT().Handle(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
		return next.Handle(res, req, nil)
	}).Times(2)

	p.Pipe(mockHandle1)
	p.Pipe(mockHandle2)

	_, err := p.Handle(mockResWriter, mockReq, mockHandleDefault)
	assert.NoError(t, err)

	_, err = p.Handle(mockResWriter, mockReq, mockHandleDefault)
	assert.NoError(t, err)
}

func TestPipeline_PipelineInPipeline(t *testing.T) {
	p1 := pipeline.New()
	p2 := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResWriter := mock_response.NewMockResponseWriter(ctrl)
	mockReq := mock_request.NewMockRequest(ctrl)

	mockHandleDefault := mock_pipeline.NewMockHandle(ctrl)
	mockHandleDefault.EXPECT().Handle(mockResWriter, mockReq, nil)

	mockHandle2 := mock_pipeline.NewMockHandle(ctrl)
	mockHandle2.EXPECT().Handle(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
		return next.Handle(res, req, nil)
	}).Times(2)

	mockHandle1 := mock_pipeline.NewMockHandle(ctrl)
	mockHandle1.EXPECT().Handle(mockResWriter, mockReq, gomock.Any()).DoAndReturn(func(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
		return next.Handle(res, req, nil)
	}).Times(2)

	p1.Pipe(mockHandle1)
	p1.Pipe(mockHandle2)

	p2.Pipe(mockHandle1)
	p2.Pipe(mockHandle2)

	p1.Pipe(p2)
	_, err := p1.Handle(mockResWriter, mockReq, mockHandleDefault)
	assert.NoError(t, err)
}

type mockMiddleware1 struct{}

func (m mockMiddleware1) Handle(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
	res.AddHeader(map[string]string{"header1": "value1"})
	return next.Handle(res, req, nil)
}

type mockMiddleware2 struct{}

func (m mockMiddleware2) Handle(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
	res.AddHeader(map[string]string{"header2": "value2"})
	return next.Handle(res, req, nil)
}

type mockMiddleware3 struct{}

func (m mockMiddleware3) Handle(res response.ResponseWriter, req request.Request, next pipeline.Handle) (response.Response, error) {
	res.AddHeader(map[string]string{"header3": "value3"})
	res.SetContent("suka")
	return res.Response(), nil
}
