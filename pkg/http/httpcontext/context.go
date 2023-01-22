package httpcontext

import (
	"net/http"

	"github.com/supermetrolog/goblack"
)

type Context struct {
	r *http.Request
	w goblack.Writer
	p map[string]string
}

func New(r *http.Request, w goblack.Writer, p map[string]string) *Context {
	return &Context{
		r: r,
		w: w,
		p: p,
	}
}

func (c Context) Request() *http.Request {
	return c.r
}
func (c Context) Writer() goblack.Writer {
	return c.w
}
func (c Context) Param(key string) string {
	value, ok := c.p[key]
	if !ok {
		return ""
	}
	return value
}
