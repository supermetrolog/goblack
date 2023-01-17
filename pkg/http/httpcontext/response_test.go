package httpcontext_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/goblack/pkg/http/httpcontext"
)

func TestNewResponse(t *testing.T) {
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"server":       {"nginx"},
		"list":         {"one", "two"},
	}

	content := []byte("test content")
	statusCode := 404
	r := httpcontext.NewResponse(content, statusCode, headers)
	assert.Equal(t, statusCode, r.StatusCode())
	assert.Equal(t, content, r.Content())
	assert.Equal(t, headers, r.Headers())
}
