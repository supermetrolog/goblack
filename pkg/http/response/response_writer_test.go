package response_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/framework/pkg/http/response"
)

func TestSetContent(t *testing.T) {
	resWriter := response.NewResponseWriter()
	content := "test content"
	contentBytes, _ := json.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestSetStatusCode(t *testing.T) {
	resWriter := response.NewResponseWriter()
	resWriter.SetStatusCode(404)
	res, err := resWriter.JsonResponse()
	assert.NoError(t, err)
	assert.Equal(t, 404, res.StatusCode())
}

func TestAddHeader(t *testing.T) {
	resWriter := response.NewResponseWriter()
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
func TestJsonResponseWithStruct(t *testing.T) {
	resWriter := response.NewResponseWriter()
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
	resWriter := response.NewResponseWriter()
	content := "content"
	contentBytes, _ := xml.Marshal(content)
	resWriter.SetContent(content)
	res, err := resWriter.XmlResponse()
	assert.NoError(t, err)
	assert.Equal(t, contentBytes, res.Content())
}
func TestXmlResponseWithArray(t *testing.T) {
	resWriter := response.NewResponseWriter()
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
