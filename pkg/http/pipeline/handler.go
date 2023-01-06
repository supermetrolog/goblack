package pipeline

import (
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
)

type Handler interface {
	Handler(res response.ResponseWriter, req request.Request) (response.Response, error)
}

type Middleware interface {
	Handler(res response.ResponseWriter, req request.Request, next Handler) (response.Response, error)
}
