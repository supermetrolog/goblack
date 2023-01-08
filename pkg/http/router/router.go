package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/framework/pkg/http/app"
	"github.com/supermetrolog/framework/pkg/http/httpcontext"
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	contextInterface "github.com/supermetrolog/framework/pkg/http/interfaces/httpcontext"
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
		w.WriteHeader(res.StatusCode())
		for key, h := range res.Headers() {
			for _, value := range h {
				w.Header().Add(key, value)
			}
		}
		w.Write(res.Content())
	}
}
