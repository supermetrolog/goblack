package httpcontext

import (
	"net/http"

	"github.com/supermetrolog/goblack/pkg/http/interfaces/httpcontext"
)

type Context struct {
	r  *http.Request
	rw httpcontext.ResponseWriter
	p  map[string]string
}

func New(r *http.Request, rw httpcontext.ResponseWriter, p map[string]string) *Context {
	return &Context{
		r:  r,
		rw: rw,
		p:  p,
	}
}

func (c Context) Request() *http.Request {
	return c.r
}
func (c Context) ResponseWriter() httpcontext.ResponseWriter {
	return c.rw
}
func (c Context) Param(key string) string {
	value, ok := c.p[key]
	if !ok {
		return ""
	}
	return value
}
