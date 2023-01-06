package pipeline

import (
	"errors"

	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
	"github.com/supermetrolog/framework/pkg/queue"
)

type next struct {
	handler  Handler
	Handlers queue.Queue
}
type nextWrapper struct {
	n *next
}

func (n nextWrapper) Handler(res response.ResponseWriter, req request.Request) (response.Response, error) {
	return n.n.Next(res, req)
}
func newNext(q queue.Queue, handler Handler) next {
	return next{
		Handlers: q,
		handler:  handler,
	}
}
func (n next) Next(res response.ResponseWriter, req request.Request) (response.Response, error) {
	if n.Handlers.IsEmpty() {
		return n.handler.Handler(res, req)
	}
	current, ok := n.Handlers.Dequeue().(Middleware)
	if !ok {
		return nil, errors.New("unknown item in Handlers Queue")
	}
	return current.Handler(res, req, nextWrapper{n: &n})
}
