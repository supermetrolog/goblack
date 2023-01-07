package request

import "io"

type Request struct {
	readCloser  io.ReadCloser
	headers     map[string]string
	queryParams map[string]string
	pathParams  map[string]string
}

func NewRequest(
	readCloser io.ReadCloser,
	headers map[string]string,
	queryParams map[string]string,
	pathParams map[string]string,
) *Request {
	return &Request{
		readCloser:  readCloser,
		headers:     headers,
		queryParams: queryParams,
		pathParams:  pathParams,
	}
}

func (r Request) Body() io.ReadCloser {
	return r.readCloser
}

func (r Request) Header(name string) string {
	value, ok := r.headers[name]
	if !ok {
		return ""
	}
	return value
}

func (r Request) Headers() map[string]string {
	return r.headers
}

func (r Request) QueryParam(name string) string {
	value, ok := r.queryParams[name]
	if !ok {
		return ""
	}
	return value
}

func (r Request) QueryParams() map[string]string {
	return r.queryParams
}

func (r Request) Param(name string) string {
	value, ok := r.pathParams[name]
	if !ok {
		return ""
	}
	return value
}
