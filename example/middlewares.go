package example

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/supermetrolog/goblack"
)

type ProfileMiddleware struct{}

func (m ProfileMiddleware) Handler(c goblack.Context, next goblack.Handler) (goblack.Response, error) {
	startTime := time.Now().UnixMicro()
	nextRes, err := next.Handler(c)
	endTime := time.Now().UnixMicro()
	delay := endTime - startTime
	delayInSeconds := float64(delay) / float64(1000000)
	c.Writer().WriteHeader("X-Profile-Time", fmt.Sprintf("%f", delayInSeconds))
	return nextRes, err
}

type Error struct {
	Err        error  `json:"err,omitempty"`
	Message    string `json:"message,omitempty"`
	Code       int    `json:"code,omitempty"`
	StatusText string `json:"status_text,omitempty"`
}

type ErrorMiddleware struct{}

func (m ErrorMiddleware) Handler(c goblack.Context, next goblack.Handler) (goblack.Response, error) {
	log.Println("ERROR MIDDLEWARE")
	nextRes, err := next.Handler(c)
	if err == nil {
		return nextRes, err
	}
	log.Println("ERROR MIDDLEWARE 2")

	c.Writer().WriteStatus(http.StatusInternalServerError)
	c.Writer().Write(Error{
		Err:        err,
		Message:    fmt.Sprintf("Middleware catch error: %v", err),
		Code:       http.StatusInternalServerError,
		StatusText: http.StatusText(http.StatusInternalServerError),
	})
	res, resErr := c.Writer().JSON()
	if resErr != nil {
		return nil, resErr
	}
	log.Println("ERROR MIDDLEWARE 3")
	log.Println(string(res.Content()))
	return res, nil
}
