package pipeline

import (
	"errors"

	"github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/pkg/queue"
)

type Pipeline struct {
	Handlers queue.Queue
}

func New() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Pipe(middleware goblack.Middleware) {
	p.Handlers.Enqueue(middleware)
}
func (p *Pipeline) Handler(c goblack.Context, handler goblack.Handler) (goblack.Response, error) {
	if handler == nil {
		return nil, errors.New("handler can not be nil")
	}
	n := newNext(p.Handlers, handler)
	return n.Next(c)
}
