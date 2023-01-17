package pipeline

import (
	"errors"

	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
	"github.com/supermetrolog/goblack/pkg/queue"
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
func (p *Pipeline) Handler(c httpcontext.Context, handler handler.Handler) (httpcontext.Response, error) {
	if handler == nil {
		return nil, errors.New("handler can not be nil")
	}
	n := newNext(p.Handlers, handler)
	return n.Next(c)
}
