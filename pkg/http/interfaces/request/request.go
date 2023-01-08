package request

import "io"

type Request interface {
	Body() io.ReadCloser
	Header(name string) string
	HeaderValues(name string) []string
	Headers() map[string][]string
	Param(name string) string
	QueryParam(name string) string
	QueryParamValues(name string) []string
	QueryParams() map[string][]string
}
