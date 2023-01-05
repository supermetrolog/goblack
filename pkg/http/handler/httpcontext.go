package handler

import "io"

type User interface {
	Identity() any
	Token() string
	IsAuth() bool
	IsGuest() bool
	SetUser(any)
}
type Http interface {
	Write(content []byte) (int, error)
	WriteHeader(status int)
	Body() io.ReadCloser
	Header(name string) string
	Param(name string) string
}
type HttpContext interface {
	User() User
	Http() Http
}
