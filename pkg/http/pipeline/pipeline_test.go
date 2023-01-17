package pipeline_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
	"github.com/supermetrolog/goblack/pkg/http/pipeline"
	mock_handler "github.com/supermetrolog/goblack/tests/mocks/pkg/http/interfaces/handler"
	mock_httpcontex "github.com/supermetrolog/goblack/tests/mocks/pkg/http/interfaces/httpcontext"
)

func TestPipeline_pipe(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mock_handler.NewMockMiddleware(ctrl)
	mockHandler2 := mock_handler.NewMockMiddleware(ctrl)

	p.Pipe(mockHandler)
	p.Pipe(mockHandler2)

	assert.NotEmpty(t, p.Handlers)
	assert.Equal(t, 2, p.Handlers.Length())
}

func TestPipeline_runWithOnlyOneHandler(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCtx := mock_httpcontex.NewMockContext(ctrl)

	mockHandler := mock_handler.NewMockHandler(ctrl)
	mockHandler.EXPECT().Handler(mockCtx)

	_, err := p.Handler(mockCtx, mockHandler)

	assert.NoError(t, err)
}

func TestPipeline_runWithNilHandler(t *testing.T) {
	p := pipeline.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCtx := mock_httpcontex.NewMockContext(ctrl)
	_, err := p.Handler(mockCtx, nil)
	assert.Error(t, err)
}

func TestPipeline_runWithManyHandlers(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCtx := mock_httpcontex.NewMockContext(ctrl)
	mockResWriter := mock_httpcontex.NewMockResponseWriter(ctrl)
	mockCtx.EXPECT().ResponseWriter().Return(mockResWriter).Times(5)
	firstCall := mockResWriter.EXPECT().AddHeader("header1", "value1")
	secondCall := mockResWriter.EXPECT().AddHeader("header2", "value2")
	thirdCall := mockResWriter.EXPECT().AddHeader("header4", "value4")
	gomock.InOrder(
		firstCall,
		secondCall,
		thirdCall,
	)

	mockResWriter.EXPECT().SetContent("content")
	mockResWriter.EXPECT().JsonResponse()

	mock1 := mockMiddleware1{}
	mock2 := mockMiddleware2{}
	last := mockHandler{}

	p.Pipe(mock1)
	p.Pipe(mock2)
	_, err := p.Handler(mockCtx, last)
	assert.NoError(t, err)
}

func TestPipeline_doubleRun(t *testing.T) {
	p := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCtx := mock_httpcontex.NewMockContext(ctrl)

	mockHandlerDefault := mock_handler.NewMockHandler(ctrl)
	mockHandlerDefault.EXPECT().Handler(mockCtx).Times(2)

	mockHandler2 := mock_handler.NewMockMiddleware(ctrl)
	mockHandler2.EXPECT().Handler(mockCtx, gomock.Any()).DoAndReturn(func(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
		return next.Handler(c)
	}).Times(2)

	mockHandler1 := mock_handler.NewMockMiddleware(ctrl)
	mockHandler1.EXPECT().Handler(mockCtx, gomock.Any()).DoAndReturn(func(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
		return next.Handler(c)
	}).Times(2)

	p.Pipe(mockHandler1)
	p.Pipe(mockHandler2)

	_, err := p.Handler(mockCtx, mockHandlerDefault)
	assert.NoError(t, err)

	_, err = p.Handler(mockCtx, mockHandlerDefault)
	assert.NoError(t, err)
}

func TestPipeline_PipelineInPipeline(t *testing.T) {
	p1 := pipeline.New()
	p2 := pipeline.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCtx := mock_httpcontex.NewMockContext(ctrl)

	mockHandlerDefault := mock_handler.NewMockHandler(ctrl)
	mockHandlerDefault.EXPECT().Handler(mockCtx)

	mockHandler2 := mock_handler.NewMockMiddleware(ctrl)
	mockHandler2.EXPECT().Handler(mockCtx, gomock.Any()).DoAndReturn(func(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
		return next.Handler(c)
	}).Times(2)

	mockHandler1 := mock_handler.NewMockMiddleware(ctrl)
	mockHandler1.EXPECT().Handler(mockCtx, gomock.Any()).DoAndReturn(func(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
		return next.Handler(c)
	}).Times(2)

	p1.Pipe(mockHandler1)
	p1.Pipe(mockHandler2)

	p2.Pipe(mockHandler1)
	p2.Pipe(mockHandler2)

	p1.Pipe(p2)
	_, err := p1.Handler(mockCtx, mockHandlerDefault)
	assert.NoError(t, err)
}

type mockMiddleware1 struct{}

func (m mockMiddleware1) Handler(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
	c.ResponseWriter().AddHeader("header1", "value1")
	return next.Handler(c)
}

type mockMiddleware2 struct{}

func (m mockMiddleware2) Handler(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
	c.ResponseWriter().AddHeader("header2", "value2")
	return next.Handler(c)
}

type mockMiddleware3 struct{}

func (m mockMiddleware3) Handler(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
	c.ResponseWriter().AddHeader("header3", "value3")
	c.ResponseWriter().SetContent("suka")
	return c.ResponseWriter().JsonResponse()
}

type mockHandler struct{}

func (m mockHandler) Handler(c httpcontext.Context) (httpcontext.Response, error) {
	c.ResponseWriter().AddHeader("header4", "value4")
	c.ResponseWriter().SetContent("content")
	return c.ResponseWriter().JsonResponse()
}
