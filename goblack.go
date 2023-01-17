package goblack

import (
	"net/http"

	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
)

type ServerConfig struct {
	Addr string
}

type Pipeline interface {
	handler.Middleware
	Pipe(handler.Middleware)
}

type Router interface {
	GET(path string, handler handler.Handler, middlewares ...handler.Middleware)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type App struct {
	serverConfig ServerConfig
	pipeline     Pipeline
	router       Router
}

func New(pipeline Pipeline, router Router, serverConfig ServerConfig) *App {
	return &App{
		pipeline:     pipeline,
		router:       router,
		serverConfig: serverConfig,
	}
}

func (app *App) Run() error {
	return http.ListenAndServe(app.serverConfig.Addr, app.router)
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

// func Default() *App {
// 	pFactory := pipeline.NewFactory()
// 	p := pFactory.Create()
// 	er := httprouter.New()
// 	r := router.New(p, pFactory, er)
// 	return New(p, r, ServerConfig{Addr: ":80"})
// }
