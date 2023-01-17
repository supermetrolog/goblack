package pipeline

import "github.com/supermetrolog/goblack"

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}
func (f Factory) Create() goblack.Pipeline {
	return New()
}
