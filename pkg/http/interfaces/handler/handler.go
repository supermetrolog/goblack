package handler

import (
	"github.com/supermetrolog/framework/pkg/http/interfaces/httpcontext"
)

type Handler interface {
	Handler(c httpcontext.Context) (httpcontext.Response, error)
}

type Middleware interface {
	Handler(c httpcontext.Context, next Handler) (httpcontext.Response, error)
}
