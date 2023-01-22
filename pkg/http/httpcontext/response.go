package httpcontext

import (
	"encoding/json"
	"encoding/xml"

	"github.com/supermetrolog/goblack"
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

type Writer struct {
	content    any
	statusCode int
	headers    map[string][]string
}

func NewWriter() *Writer {
	return &Writer{
		headers: make(map[string][]string),
	}
}
func (r *Writer) Write(content any) goblack.Writer {
	r.content = content
	return r
}
func (r *Writer) WriteStatus(statusCode int) goblack.Writer {
	r.statusCode = statusCode
	return r
}
func (r *Writer) WriteHeader(key string, value string) goblack.Writer {
	r.headers[key] = append(r.headers[key], value)
	return r
}
func (r *Writer) HasHeaderValue(key string, value string) bool {
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
func (r Writer) JSON() (goblack.Response, error) {
	bytes, err := json.Marshal(r.content)
	if err != nil {
		return nil, err
	}

	if !r.HasHeaderValue("Content-Type", "application/json") {
		r.WriteHeader("Content-Type", "application/json")
	}
	response := NewResponse(bytes, r.statusCode, r.headers)
	return response, nil
}
func (r Writer) HTML() (goblack.Response, error) {
	return nil, nil
}
func (r Writer) XML() (goblack.Response, error) {
	bytes, err := xml.Marshal(r.content)
	if err != nil {
		return nil, err
	}
	response := NewResponse(bytes, r.statusCode, r.headers)
	return response, nil
}
