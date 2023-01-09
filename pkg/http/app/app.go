package app

import (
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	"github.com/supermetrolog/framework/pkg/http/interfaces/httpcontext"
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
	router   Router
}

func New(pipeline Pipeline, router Router) *App {
	return &App{
		pipeline: pipeline,
		router:   router,
	}
}

func (app *App) Pipe(middleware handler.Middleware) {
	app.pipeline.Pipe(middleware)
}

func (app App) Handler(c httpcontext.Context, next handler.Handler) (httpcontext.Response, error) {
	return app.pipeline.Handler(c, next)
}

func (app *App) GET(path string, handler handler.Handler, middlewares ...handler.Middleware) {
	app.router.GET(path, handler, middlewares...)
}
