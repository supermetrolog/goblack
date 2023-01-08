package httpcontext

import (
	"net/http"
)

type Response interface {
	Content() []byte
	StatusCode() int
	Headers() map[string][]string
}

type ResponseWriter interface {
	SetContent(any) ResponseWriter
	SetStatusCode(int) ResponseWriter
	AddHeader(key string, value string) ResponseWriter
	JsonResponse() (Response, error)
	HtmlResponse() (Response, error)
	XmlResponse() (Response, error)
}

type Context interface {
	Request() *http.Request
	Param(key string) string
	ResponseWriter() ResponseWriter
}
