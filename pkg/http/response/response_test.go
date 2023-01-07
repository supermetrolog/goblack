package response_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/framework/pkg/http/response"
)

func TestNewResponse(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"server":       "nginx",
	}

	content := []byte("test content")
	statusCode := 404
	r := response.NewResponse(content, statusCode, headers)
	assert.Equal(t, statusCode, r.StatusCode())
	assert.Equal(t, content, r.Content())
	assert.Equal(t, headers, r.Headers())
}
