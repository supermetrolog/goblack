package router

import "github.com/supermetrolog/framework/pkg/http/interfaces/handler"

func (router Router) GET(path string, handler handler.Handler, middlewares ...handler.Middleware) {
	router.externalRouter.GET(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
