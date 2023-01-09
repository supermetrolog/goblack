package httpcontext

import (
	"encoding/json"
	"encoding/xml"

	"github.com/supermetrolog/framework/pkg/http/interfaces/httpcontext"
)

type Response struct {
	content    []byte
	statusCode int
	headers    map[string][]string
}

func NewResponse(
	content []byte,
	statusCode int,
	headers map[string][]string,
) *Response {
	return &Response{
		content:    content,
		statusCode: statusCode,
		headers:    headers,
	}
}
func (r Response) Content() []byte {
	return r.content
}
func (r Response) StatusCode() int {
	return r.statusCode
}
func (r Response) Headers() map[string][]string {
	return r.headers
}

type ResponseWriter struct {
	content    any
	statusCode int
	headers    map[string][]string
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		headers: make(map[string][]string),
	}
}
func (r *ResponseWriter) SetContent(content any) httpcontext.ResponseWriter {
	r.content = content
	return r
}
func (r *ResponseWriter) SetStatusCode(statusCode int) httpcontext.ResponseWriter {
	r.statusCode = statusCode
	return r
}
func (r *ResponseWriter) AddHeader(key string, value string) httpcontext.ResponseWriter {
	r.headers[key] = append(r.headers[key], value)
	return r
}
func (r *ResponseWriter) HasHeaderValue(key string, value string) bool {
	header, ok := r.headers[key]
	if !ok {
		return false
	}
	for _, v := range header {
		if v == value {
			return true
		}
	}
	return false
}
func (r ResponseWriter) JsonResponse() (httpcontext.Response, error) {
	bytes, err := json.Marshal(r.content)
	if err != nil {
		return nil, err
	}

	if !r.HasHeaderValue("Content-Type", "application/json") {
		r.AddHeader("Content-Type", "application/json")
	}
	response := NewResponse(bytes, r.statusCode, r.headers)
	return response, nil
}
func (r ResponseWriter) HtmlResponse() (httpcontext.Response, error) {
	return nil, nil
}
func (r ResponseWriter) XmlResponse() (httpcontext.Response, error) {
	bytes, err := xml.Marshal(r.content)
	if err != nil {
		return nil, err
	}
	response := NewResponse(bytes, r.statusCode, r.headers)
	return response, nil
}
