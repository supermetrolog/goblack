package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/supermetrolog/goblack"
	"github.com/supermetrolog/goblack/example"
	"github.com/supermetrolog/goblack/pkg/http/pipeline"
	"github.com/supermetrolog/goblack/pkg/http/router"
)

func main() {
	fmt.Println("MAIN")
	pipelineFactory := pipeline.NewFactory()
	pipelineMain := pipelineFactory.Create()
	externalRouter := httprouter.New()
	app := goblack.New(pipelineMain, router.New(pipelineMain, pipelineFactory, externalRouter), goblack.ServerConfig{
		Addr: ":8080",
	})
	app.Pipe(example.ProfileMiddleware{})
	app.Pipe(example.ErrorMiddleware{})
	app.GET("/users", example.UserListHandler{})
	app.GET("/users/:id", example.UserHandler{})
	// http.ListenAndServe(":8080", externalRouter)
	app.Run()
}
