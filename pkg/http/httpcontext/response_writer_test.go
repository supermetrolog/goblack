package httpcontext_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermetrolog/goblack/pkg/http/httpcontext"
)

func TestSetContent(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := "test content"
	contentBytes, _ := json.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestSetStatusCode(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	resWriter.SetStatusCode(404)
	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, 404, res.StatusCode())
}

func TestAddHeader(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"server":       {"nginx"},
		"allow":        {"*"},
		"list":         {"one", "two"},
	}
	for key, h := range headers {
		for _, value := range h {
			resWriter.AddHeader(key, value)

		}
	}

	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, headers, res.Headers())
}
func TestJsonResponse(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := "content"
	contentBytes, _ := json.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.JsonResponse()
	contentTypes, ok := res.Headers()["Content-Type"]
	require.True(t, ok)
	assert.Equal(t, "application/json", contentTypes[0])
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestJsonResponseDoubleCall(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := "content"
	contentBytes, _ := json.Marshal(content)
	resWriter.SetContent(content)
	resWriter.JsonResponse()
	res, err := resWriter.JsonResponse()
	contentTypes, ok := res.Headers()["Content-Type"]
	require.True(t, ok)
	assert.Len(t, contentTypes, 1)
	assert.Equal(t, "application/json", contentTypes[0])
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestJsonResponseWithStruct(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := struct {
		Username string
		Password string
		Name     string
	}{
		Username: "John",
		Password: "qwerty",
		Name:     "Dodson",
	}
	contentBytes, _ := json.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestXmlResponseWithString(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := "content"
	contentBytes, _ := xml.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.XmlResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestXmlResponseWithArray(t *testing.T) {
	resWriter := httpcontext.NewResponseWriter()
	content := []string{
		"John",
		"qwerty",
		"Dodson",
	}
	contentBytes, _ := xml.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.XmlResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
