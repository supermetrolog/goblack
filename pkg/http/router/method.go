package router

import (
	"net/http"

	"github.com/supermetrolog/goblack"
)

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.externalRouter.ServeHTTP(w, r)
}

func (router Router) GET(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.GET(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
func (router Router) POST(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.POST(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}

func (router Router) PUT(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.PUT(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}

func (router Router) PATCH(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.PATCH(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
func (router Router) DELETE(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.DELETE(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
func (router Router) OPTIONS(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.OPTIONS(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
func (router Router) HEAD(path string, handler goblack.Handler, middlewares ...goblack.Middleware) {
	router.externalRouter.HEAD(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
