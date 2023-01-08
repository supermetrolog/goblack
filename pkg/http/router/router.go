package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/framework/pkg/http/app"
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	reqst "github.com/supermetrolog/framework/pkg/http/request"
	"github.com/supermetrolog/framework/pkg/http/response"
)

type PipelineFactory interface {
	Create() app.Pipeline
}

type Router struct {
	mainPipeline    app.Pipeline
	pipelineFactory PipelineFactory
	externalRouter  *httprouter.Router
}

func NewRouter(mainPipeline app.Pipeline, pipelineFactory PipelineFactory) *Router {
	return &Router{
		mainPipeline:    mainPipeline,
		pipelineFactory: pipelineFactory,
	}
}
func (r Router) makePipeline(middlewares []handler.Middleware) app.Pipeline {
	pipeline := r.pipelineFactory.Create()
	pipeline.Pipe(r.mainPipeline)
	for _, middleware := range middlewares {
		pipeline.Pipe(middleware)
	}
	return pipeline
}
func (r Router) makeRequest(defaultRequest *http.Request, params httprouter.Params) request.Request {
	paramsMap := make(map[string]string, len(params))
	for _, p := range params {
		paramsMap[p.Key] = p.Value
	}
	return reqst.NewRequest(defaultRequest, paramsMap)
}
func (router Router) makeHandlerAdapter(handler handler.Handler, middlewares []handler.Middleware) httprouter.Handle {
	pipeline := router.makePipeline(middlewares)
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := router.makeRequest(r, p)
		resWriter := response.NewResponseWriter()
		res, err := pipeline.Handler(resWriter, req, handler)
		if err != nil {
			return
		}
		w.WriteHeader(res.StatusCode())
		for key, h := range res.Headers() {
			for _, value := range h {
				w.Header().Add(key, value)
			}
		}
		w.Write(res.Content())
	}
}
