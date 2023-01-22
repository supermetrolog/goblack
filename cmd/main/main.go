package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/pkg/http/pipeline"
	"github.com/supermetrolog/goblack/pkg/http/router"
)

type LoggerMiddleware struct{}

func (l LoggerMiddleware) Handler(c goblack.Context, next goblack.Handler) (goblack.Response, error) {
	startTime := time.Now().UnixMicro()
	nextRes, err := next.Handler(c)
	endTime := time.Now().UnixMicro()
	delay := endTime - startTime
	delayInSeconds := float64(delay) / float64(1000000)
	c.Writer().WriteHeader("X-Profile-Time", fmt.Sprintf("%f", delayInSeconds))
	return nextRes, err
}

type LoggerMiddleware2 struct{}

func (l LoggerMiddleware2) Handler(c goblack.Context, next goblack.Handler) (goblack.Response, error) {
	log.Println("Logger middleware2")
	c.Writer().Write("fuck")
	next.Handler(c)
	c.Writer().WriteHeader("fuck", "suck")
	return c.Writer().JSON()
}

type Handler struct {
	logger log.Logger
}

func NewHandler(logger log.Logger) Handler {
	return Handler{
		logger: logger,
	}
}
func (l Handler) Handler(c goblack.Context) (goblack.Response, error) {
	log.Println("Handler")
	array := []string{"gomosek", "4mo"}
	c.Writer().WriteStatus(http.StatusBadRequest)
	c.Writer().Write(array)
	c.Writer().WriteHeader("nigga", "pidor")
	return c.Writer().JSON()
}

type UserHandler struct {
	logger log.Logger
}

func NewUserHandler(logger log.Logger) UserHandler {
	return UserHandler{
		logger: logger,
	}
}
func (uh UserHandler) Handler(c goblack.Context) (goblack.Response, error) {
	log.Println("NIGGA")
	id := c.Param("id")
	c.Writer().WriteStatus(http.StatusOK)
	c.Writer().Write(id)
	return c.Writer().JSON()
}
func main() {
	fmt.Println("MAIN")
	pipelineFactory := pipeline.NewFactory()
	pipelineMain := pipelineFactory.Create()
	externalRouter := httprouter.New()
	app := goblack.New(pipelineMain, router.New(pipelineMain, pipelineFactory, externalRouter), goblack.ServerConfig{
		Addr: ":8080",
	})
	app.Pipe(LoggerMiddleware{})
	app.GET("/users", NewHandler(log.Logger{}), LoggerMiddleware2{})
	app.GET("/users/:id", NewUserHandler(log.Logger{}))
	// http.ListenAndServe(":8080", externalRouter)
	app.Run()
}
