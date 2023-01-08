package main

import (
	"fmt"
	"log"
	"net/http"

	application "github.com/supermetrolog/framework/pkg/http/app"
	"github.com/supermetrolog/framework/pkg/http/httpcontext"
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	httpcontextInterface "github.com/supermetrolog/framework/pkg/http/interfaces/httpcontext"
	"github.com/supermetrolog/framework/pkg/http/pipeline"
)

type AdapterMiddleware struct {
	W http.ResponseWriter
	R *http.Request
}

func (l AdapterMiddleware) Handler(c httpcontextInterface.Context, next handler.Handler) (httpcontextInterface.Response, error) {
	log.Println("Logger middleware")

	nextRes, err := next.Handler(c)
	log.Println(nextRes.Headers())
	return nextRes, err
}

type LoggerMiddleware struct{}

func (l LoggerMiddleware) Handler(c httpcontextInterface.Context, next handler.Handler) (httpcontextInterface.Response, error) {
	log.Println("Logger middleware")

	nextRes, err := next.Handler(c)
	log.Println(nextRes.Headers())
	return nextRes, err
}

type LoggerMiddleware2 struct{}

func (l LoggerMiddleware2) Handler(c httpcontextInterface.Context, next handler.Handler) (httpcontextInterface.Response, error) {
	log.Println("Logger middleware2")
	c.ResponseWriter().SetContent("fuck")
	next.Handler(c)
	c.ResponseWriter().SetContent("fuck")
	c.ResponseWriter().AddHeader("fuck", "suck")
	return c.ResponseWriter().JsonResponse()
}

type Handler struct {
	logger log.Logger
}

func NewHandler(logger log.Logger) Handler {
	return Handler{
		logger: logger,
	}
}
func (l Handler) Handler(c httpcontextInterface.Context) (httpcontextInterface.Response, error) {
	log.Println("Handler")
	array := []string{"nigger", "fuck", "suck"}
	c.ResponseWriter().SetContent(array)
	c.ResponseWriter().AddHeader("nigga", "pidor")
	return c.ResponseWriter().JsonResponse()
}
func main() {
	fmt.Println("MAIN")
	r, _ := http.NewRequest("GET", "/users", nil)
	httpContext := httpcontext.New(r, httpcontext.NewResponseWriter(), map[string]string{"id": "12", "test": "1234"})

	app := application.New(pipeline.New())
	app.Pipe(LoggerMiddleware{})
	app.Pipe(LoggerMiddleware{})
	app.Pipe(LoggerMiddleware2{})

	app2 := application.New(pipeline.New())
	app2.Pipe(LoggerMiddleware{})
	app.Pipe(app2)
	app.Handler(httpContext, NewHandler(log.Logger{}))
	app.GET("/users", NewHandler(log.Logger{}), LoggerMiddleware{}, LoggerMiddleware2{})
}
