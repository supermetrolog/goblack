package app

import (
	"github.com/supermetrolog/framework/pkg/http/pipeline"
)

type App struct {
	pipeline.Pipeline
}

func New() *App {
	return &App{}
}
