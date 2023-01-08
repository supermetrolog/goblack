package request_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/framework/pkg/http/request"
)

var (
	headers = map[string][]string{
		"Content-Type": {"application/json"},
		"Refer":        {"https://google.com/"},
		"Pragma":       {"no-cache"},
		"List":         {"one", "two", "three"},
	}

	queryParams = map[string][]string{
		"expand": {"files", "orders", "profile"},
		"token":  {"asdadwdasdw"},
	}

	pathParams = map[string]string{
		"id": "213",
	}
)

func CreateRequest() *request.Request {
	r := httptest.NewRequest("GET", "/path", nil)
	for key, h := range headers {
		for _, value := range h {
			r.Header.Add(key, value)
		}
	}
	query := r.URL.Query()
	for key, q := range queryParams {
		for _, value := range q {
			query.Add(key, value)
		}
	}
	r.URL.RawQuery = query.Encode()
	return request.NewRequest(r, pathParams)
}

func TestHeader(t *testing.T) {
	r := CreateRequest()
	refer := r.Header("Refer")
	contentType := r.Header("Content-Type")
	assert.Equal(t, headers["Refer"][0], refer)
	assert.Equal(t, headers["Content-Type"][0], contentType)
}

func TestHeaders(t *testing.T) {
	r := CreateRequest()
	actualHeaders := r.Headers()
	assert.Equal(t, headers, actualHeaders)
}

func TestQueryParam(t *testing.T) {
	r := CreateRequest()
	expand := r.QueryParam("expand")
	token := r.QueryParam("token")
	assert.Equal(t, queryParams["expand"][0], expand)
	assert.Equal(t, queryParams["token"][0], token)
}
func TestQueryParamValues(t *testing.T) {
	r := CreateRequest()
	expand := r.QueryParamValues("expand")
	token := r.QueryParamValues("token")
	assert.Equal(t, queryParams["expand"], expand)
	assert.Equal(t, queryParams["token"], token)
}
func TestQueryParams(t *testing.T) {
	r := CreateRequest()
	actualQueryParams := r.QueryParams()
	assert.Equal(t, queryParams, actualQueryParams)
}
func TestHeaderValues(t *testing.T) {
	r := CreateRequest()
	values := r.HeaderValues("List")
	assert.Equal(t, headers["List"], values)
}
func TestParam(t *testing.T) {
	r := CreateRequest()
	id := r.Param("id")
	assert.Equal(t, pathParams["id"], id)
}
