package pipeline

import (
	"errors"

	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
	"github.com/supermetrolog/goblack/pkg/queue"
)

type next struct {
	handler  handler.Handler
	Handlers queue.Queue
}
type nextWrapper struct {
	n *next
}

func (n nextWrapper) Handler(c httpcontext.Context) (httpcontext.Response, error) {
	return n.n.Next(c)
}
func newNext(q queue.Queue, handler handler.Handler) next {
	return next{
		Handlers: q,
		handler:  handler,
	}
}
func (n next) Next(c httpcontext.Context) (httpcontext.Response, error) {
	if n.Handlers.IsEmpty() {
		return n.handler.Handler(c)
	}
	current, ok := n.Handlers.Dequeue().(handler.Middleware)
	if !ok {
		return nil, errors.New("unknown item in Handlers Queue")
	}
	return current.Handler(c, nextWrapper{n: &n})
}
