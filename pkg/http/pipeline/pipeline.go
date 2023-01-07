package pipeline

import (
	"errors"

	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
	"github.com/supermetrolog/framework/pkg/queue"
)

type Pipeline struct {
	Handlers queue.Queue
}

func New() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Pipe(middleware handler.Middleware) {
	p.Handlers.Enqueue(middleware)
}
func (p *Pipeline) Handler(res response.ResponseWriter, req request.Request, handler handler.Handler) (response.Response, error) {
	if handler == nil {
		return nil, errors.New("handler can not be nil")
	}
	n := newNext(p.Handlers, handler)
	return n.Next(res, req)
}
