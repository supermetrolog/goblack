package pipeline

import (
	"errors"

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

func (p *Pipeline) Pipe(handle Handle) {
	p.Handlers.Enqueue(handle)
}
func (p *Pipeline) Handle(res response.ResponseWriter, req request.Request, nextDefault Handle) (response.Response, error) {
	if nextDefault == nil {
		return nil, errors.New("default Handle can not be nil")
	}
	n := newNext(p.Handlers, nextDefault)
	return n.Next(res, req)
}
