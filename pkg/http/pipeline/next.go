package pipeline

import (
	"errors"

	"github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/pkg/queue"
)

type next struct {
	handler  goblack.Handler
	Handlers queue.Queue
}
type nextWrapper struct {
	n *next
}

func (n nextWrapper) Handler(c goblack.Context) (goblack.Response, error) {
	return n.n.Next(c)
}
func newNext(q queue.Queue, handler goblack.Handler) next {
	return next{
		Handlers: q,
		handler:  handler,
	}
}
func (n next) Next(c goblack.Context) (goblack.Response, error) {
	if n.Handlers.IsEmpty() {
		return n.handler.Handler(c)
	}
	current, ok := n.Handlers.Dequeue().(goblack.Middleware)
	if !ok {
		return nil, errors.New("unknown item in Handlers Queue")
	}
	return current.Handler(c, nextWrapper{n: &n})
}
