package pipeline

import (
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
)

type Handle interface {
	Handle(res response.ResponseWriter, req request.Request, next Handle) (response.Response, error)
}
