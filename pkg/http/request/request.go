package request

import (
	"io"
	"net/http"
)

type Request struct {
	r *http.Request
	p map[string]string
}

func NewRequest(r *http.Request, p map[string]string) *Request {
	return &Request{
		r: r,
		p: p,
	}
}

func (r Request) Body() io.ReadCloser {
	return r.r.Body
}

func (r Request) Header(name string) string {
	return r.r.Header.Get(name)
}
func (r Request) HeaderValues(name string) []string {
	return r.r.Header.Values(name)
}
func (r Request) Headers() map[string][]string {
	return r.r.Header
}

func (r Request) QueryParam(name string) string {
	return r.r.URL.Query().Get(name)
}
func (r Request) QueryParamValues(name string) []string {
	q := r.r.URL.Query()
	values, ok := q[name]
	if !ok {
		return nil
	}
	return values
}
func (r Request) QueryParams() map[string][]string {
	return r.r.URL.Query()
}

func (r Request) Param(name string) string {
	value, ok := r.p[name]
	if !ok {
		return ""
	}
	return value
}
