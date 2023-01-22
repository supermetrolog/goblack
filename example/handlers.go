package example

import (
	"errors"
	"log"
	"net/http"

	"github.com/supermetrolog/goblack"
)

type UserListHandler struct{}

func (h UserListHandler) Handler(c goblack.Context) (goblack.Response, error) {
	log.Println("Handler")
	array := []string{"gomosek", "4mo"}
	c.Writer().WriteStatus(http.StatusBadRequest)
	c.Writer().Write(array)
	c.Writer().WriteHeader("nigga", "pidor")
	return c.Writer().JSON()
}

type UserHandler struct{}

func (uh UserHandler) Handler(c goblack.Context) (goblack.Response, error) {
	id := c.Param("id")
	c.Writer().WriteStatus(http.StatusOK)
	c.Writer().Write(id)

	res, err := c.Writer().JSON()
	if err != nil {
		return nil, err
	}
	return res, errors.New("user handler error")
}
