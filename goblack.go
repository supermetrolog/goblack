package goblack

import (
	"net/http"
)

//go:generate mockgen -destination=mocks/mock_goblack.go -package=mock_goblack . Pipeline,Router,Handler,Middleware,Context,Response,ResponseWriter
type Pipeline interface {
	Middleware
	Pipe(Middleware)
}
type Router interface {
	GET(path string, handler Handler, middlewares ...Middleware)
	POST(path string, handler Handler, middlewares ...Middleware)
	PUT(path string, handler Handler, middlewares ...Middleware)
	PATCH(path string, handler Handler, middlewares ...Middleware)
	DELETE(path string, handler Handler, middlewares ...Middleware)
	OPTIONS(path string, handler Handler, middlewares ...Middleware)
	HEAD(path string, handler Handler, middlewares ...Middleware)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
