package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	application "github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/pkg/http/interfaces/handler"
	httpcontextInterface "github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
	"github.com/supermetrolog/goblack/pkg/http/pipeline"
	"github.com/supermetrolog/goblack/pkg/http/router"
)

type LoggerMiddleware struct{}

func (l LoggerMiddleware) Handler(c httpcontextInterface.Context, next handler.Handler) (httpcontextInterface.Response, error) {
	startTime := time.Now().UnixMicro()
	nextRes, err := next.Handler(c)
	endTime := time.Now().UnixMicro()
	delay := endTime - startTime
	delayInSeconds := float64(delay) / float64(1000000)
	c.ResponseWriter().AddHeader("X-Profile-Time", fmt.Sprintf("%f", delayInSeconds))
	return nextRes, err
}

type LoggerMiddleware2 struct{}

func (l LoggerMiddleware2) Handler(c httpcontextInterface.Context, next handler.Handler) (httpcontextInterface.Response, error) {
	log.Println("Logger middleware2")
	c.ResponseWriter().SetContent("fuck")
	next.Handler(c)
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
	array := []string{"gomosek", "4mo"}
	c.ResponseWriter().SetStatusCode(http.StatusBadRequest)
	c.ResponseWriter().SetContent(array)
	c.ResponseWriter().AddHeader("nigga", "pidor")
	return c.ResponseWriter().JsonResponse()
}

type UserHandler struct {
	logger log.Logger
}

func NewUserHandler(logger log.Logger) UserHandler {
	return UserHandler{
		logger: logger,
	}
}
func (uh UserHandler) Handler(c httpcontextInterface.Context) (httpcontextInterface.Response, error) {
	log.Println("NIGGA")
	id := c.Param("id")
	c.ResponseWriter().SetStatusCode(http.StatusOK)
	c.ResponseWriter().SetContent(id)
	return c.ResponseWriter().JsonResponse()
}
func main() {
	fmt.Println("MAIN")
	pipelineFactory := pipeline.NewFactory()
	pipelineMain := pipelineFactory.Create()
	externalRouter := httprouter.New()
	app := application.New(pipelineMain, router.New(pipelineMain, pipelineFactory, externalRouter), application.ServerConfig{
		Addr: ":8080",
	})
	app.Pipe(LoggerMiddleware{})
	app.GET("/users", NewHandler(log.Logger{}), LoggerMiddleware2{})
	app.GET("/users/:id", NewUserHandler(log.Logger{}))
	// http.ListenAndServe(":8080", externalRouter)
	app.Run()
}
