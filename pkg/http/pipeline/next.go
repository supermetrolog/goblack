package pipeline

import (
	"errors"

	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
	"github.com/supermetrolog/framework/pkg/queue"
)

type next struct {
	nextDefault Handle
	Handlers    queue.Queue
}
type nextWrapper struct {
	n *next
}

func (n nextWrapper) Handle(res response.ResponseWriter, req request.Request, next Handle) (response.Response, error) {
	return n.n.Next(res, req)
}
func newNext(q queue.Queue, nextDefault Handle) next {
	return next{
		Handlers:    q,
		nextDefault: nextDefault,
	}
}
func (n next) Next(res response.ResponseWriter, req request.Request) (response.Response, error) {
	if n.Handlers.IsEmpty() {
		return n.nextDefault.Handle(res, req, nil)
	}
	current, ok := n.Handlers.Dequeue().(Handle)
	if !ok {
		return nil, errors.New("unknown item in Handlers Queue")
	}
	return current.Handle(res, req, nextWrapper{n: &n})
}
