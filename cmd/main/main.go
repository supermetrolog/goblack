package main

import (
	"fmt"
	"log"
	"net/http"

	application "github.com/supermetrolog/framework/pkg/http/app"
	"github.com/supermetrolog/framework/pkg/http/interfaces/handler"
	"github.com/supermetrolog/framework/pkg/http/interfaces/request"
	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
	"github.com/supermetrolog/framework/pkg/http/pipeline"
	reqst "github.com/supermetrolog/framework/pkg/http/request"
	resps "github.com/supermetrolog/framework/pkg/http/response"
)

type AdapterMiddleware struct {
	W http.ResponseWriter
	R *http.Request
}

func (l AdapterMiddleware) Handler(res response.ResponseWriter, req request.Request, next handler.Handler) (response.Response, error) {
	log.Println("Logger middleware")

	nextRes, err := next.Handler(res, req)
	log.Println(nextRes.Headers())
	return nextRes, err
}

type LoggerMiddleware struct{}

func (l LoggerMiddleware) Handler(res response.ResponseWriter, req request.Request, next handler.Handler) (response.Response, error) {
	log.Println("Logger middleware")

	nextRes, err := next.Handler(res, req)
	log.Println(nextRes.Headers())
	return nextRes, err
}

type LoggerMiddleware2 struct{}

func (l LoggerMiddleware2) Handler(res response.ResponseWriter, req request.Request, next handler.Handler) (response.Response, error) {
	log.Println("Logger middleware2")
	res.SetContent("fuck")
	next.Handler(res, req)
	res.SetContent("fuck")
	res.AddHeader("fuck", "suck")
	return res.JsonResponse()
}

type Handler struct {
	logger log.Logger
}

func NewHandler(logger log.Logger) Handler {
	return Handler{
		logger: logger,
	}
}
func (l Handler) Handler(res response.ResponseWriter, req request.Request) (response.Response, error) {
	log.Println("Handler")
	array := []string{"nigger", "fuck", "suck"}
	res.SetContent(array)
	res.AddHeader("nigga", "pidor")
	return res.JsonResponse()
}
func main() {
	fmt.Println("MAIN")
	app := application.New(pipeline.New())
	app.Pipe(LoggerMiddleware{})
	app.Pipe(LoggerMiddleware{})
	app.Pipe(LoggerMiddleware2{})

	app2 := application.New(pipeline.New())
	app2.Pipe(LoggerMiddleware{})
	app.Pipe(app2)
	app.Handler(resps.NewResponseWriter(), reqst.NewRequest(nil, nil, nil), NewHandler(log.Logger{}))
	// app.GET("/users", NewHandler(log.Logger{}), LoggerMiddleware, LoggerMiddleware2)
}
