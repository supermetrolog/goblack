package response

import (
	"encoding/json"

	"github.com/supermetrolog/framework/pkg/http/interfaces/response"
)

type Response struct {
	content    []byte
	statusCode int
	headers    map[string]string
}

func NewResponse(
	content []byte,
	statusCode int,
	headers map[string]string,
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
func (r Response) Headers() map[string]string {
	return r.headers
}

type ResponseWriter struct {
	content    any
	statusCode int
	headers    map[string]string
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		headers: make(map[string]string),
	}
}
func (r *ResponseWriter) SetContent(content any) {
	r.content = content
}
func (r *ResponseWriter) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}
func (r *ResponseWriter) AddHeader(key string, value string) {
	r.headers[key] = value
}
func (r ResponseWriter) JsonResponse() (response.Response, error) {
	bytes, err := json.Marshal(r.content)
	if err != nil {
		return nil, err
	}
	response := NewResponse(bytes, r.statusCode, r.headers)
	return response, nil
}
func (r ResponseWriter) HtmlResponse() (response.Response, error) {
	return nil, nil
}
func (r ResponseWriter) Response() response.Response {
	return nil
}
