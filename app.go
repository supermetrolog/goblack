package goblack

import "net/http"

type ServerConfig struct {
	Addr string
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

func (app *App) Pipe(middleware Middleware) {
	app.pipeline.Pipe(middleware)
}

func (app App) Handler(c Context, next Handler) (Response, error) {
	return app.pipeline.Handler(c, next)
}
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}
func (app *App) GET(path string, handler Handler, middlewares ...Middleware) {
	app.router.GET(path, handler, middlewares...)
}
func (app *App) POST(path string, handler Handler, middlewares ...Middleware) {
	app.router.POST(path, handler, middlewares...)
}
func (app *App) PUT(path string, handler Handler, middlewares ...Middleware) {
	app.router.PUT(path, handler, middlewares...)
}
func (app *App) PATCH(path string, handler Handler, middlewares ...Middleware) {
	app.router.PATCH(path, handler, middlewares...)
}
func (app *App) DELETE(path string, handler Handler, middlewares ...Middleware) {
	app.router.DELETE(path, handler, middlewares...)
}
func (app *App) OPTIONS(path string, handler Handler, middlewares ...Middleware) {
	app.router.OPTIONS(path, handler, middlewares...)
}
func (app *App) HEAD(path string, handler Handler, middlewares ...Middleware) {
	app.router.HEAD(path, handler, middlewares...)
}
