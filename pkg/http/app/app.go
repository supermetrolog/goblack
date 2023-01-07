package app

import (
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
)

type Pipeline interface {
	handler.Middleware
	Pipe(handler.Middleware)
}

type Router interface {
	GET(path string, handler handler.Handler, middlewares ...handler.Middleware)
}

type App struct {
	pipeline Pipeline
}

func New(pipeline Pipeline) *App {
	return &App{
		pipeline: pipeline,
	}
}

func (app *App) Pipe(middleware handler.Middleware) {
	app.pipeline.Pipe(middleware)
}

func (app App) Handler(res response.ResponseWriter, req request.Request, next handler.Handler) (response.Response, error) {
	return app.pipeline.Handler(res, req, next)
}

func (app *App) GET(path string, handler handler.Handler, middlewares ...handler.Middleware) {

}
