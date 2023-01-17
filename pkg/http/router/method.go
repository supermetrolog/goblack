package router

import (
	"net/http"

	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
)

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.externalRouter.ServeHTTP(w, r)
}

func (router Router) GET(path string, handler handler.Handler, middlewares ...handler.Middleware) {
	router.externalRouter.GET(
		path,
		router.makeHandlerAdapter(handler, middlewares),
	)
}
