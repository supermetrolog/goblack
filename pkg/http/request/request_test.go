package request_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/framework/pkg/http/request"
)

var (
	headers = map[string]string{
		"Content-Type": "application/json",
		"refer":        "https://google.com/",
		"pragma":       "no-cache",
	}

	queryParams = map[string]string{
		"expand": "files,orders,profile",
		"token":  "asdadwdasdw",
	}

	pathParams = map[string]string{
		"id": "213",
	}
)

func CreateRequest() *request.Request {

	return request.NewRequest(nil, headers, queryParams, pathParams)
}

func TestHeader(t *testing.T) {
	r := CreateRequest()
	refer := r.Header("refer")
	contentType := r.Header("Content-Type")
	assert.Equal(t, headers["refer"], refer)
	assert.Equal(t, headers["Content-Type"], contentType)
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
	assert.Equal(t, queryParams["expand"], expand)
	assert.Equal(t, queryParams["token"], token)
}

func TestQueryParams(t *testing.T) {
	r := CreateRequest()
	actualQueryParams := r.QueryParams()
	assert.Equal(t, queryParams, actualQueryParams)
}

func TestParam(t *testing.T) {
	r := CreateRequest()
	id := r.Param("id")
	assert.Equal(t, pathParams["id"], id)
}
