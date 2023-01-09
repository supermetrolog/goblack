package pipeline

import "github.com/supermetrolog/framework/pkg/http/app"

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}
func (f Factory) Create() app.Pipeline {
	return New()
}
