package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/pkg/http/httpcontext"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	contextInterface "github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
)

type PipelineFactory interface {
	Create() goblack.Pipeline
}

type Router struct {
	mainPipeline    goblack.Pipeline
	pipelineFactory PipelineFactory
	externalRouter  *httprouter.Router
}

func New(mainPipeline goblack.Pipeline, pipelineFactory PipelineFactory, externalRouter *httprouter.Router) *Router {
	return &Router{
		mainPipeline:    mainPipeline,
		pipelineFactory: pipelineFactory,
		externalRouter:  externalRouter,
	}
}
func (router Router) makePipeline(middlewares []handler.Middleware) goblack.Pipeline {
	pipeline := router.pipelineFactory.Create()
	pipeline.Pipe(router.mainPipeline)
	for _, middleware := range middlewares {
		pipeline.Pipe(middleware)
	}
	return pipeline
}
func (router Router) makeHttpContext(r *http.Request, rw contextInterface.ResponseWriter, p httprouter.Params) contextInterface.Context {
	params := make(map[string]string, len(p))
	for _, param := range p {
		params[param.Key] = param.Value
	}
	return httpcontext.New(r, rw, params)
}
func (router Router) makeHandlerAdapter(handler handler.Handler, middlewares []handler.Middleware) httprouter.Handle {
	pipeline := router.makePipeline(middlewares)
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		httpCtx := router.makeHttpContext(r, httpcontext.NewResponseWriter(), p)
		res, err := pipeline.Handler(httpCtx, handler)
		if err != nil {
			return
		}

		for key, h := range res.Headers() {
			for _, value := range h {
				w.Header().Add(key, value)
			}
		}
		if res.StatusCode() != 0 {
			w.WriteHeader(res.StatusCode())
		}
		w.Write(res.Content())
	}
}
