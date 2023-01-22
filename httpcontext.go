package goblack

import "net/http"

type Response interface {
	Content() []byte
	StatusCode() int
	Headers() map[string][]string
}

type Writer interface {
	Write(any) Writer
	WriteStatus(int) Writer
	WriteHeader(key string, value string) Writer
	JSON() (Response, error)
	HTML() (Response, error)
	XML() (Response, error)
}

type Context interface {
	Request() *http.Request
	Param(key string) string
	Writer() Writer
}
