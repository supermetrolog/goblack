package httpcontext_test

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supermetrolog/goblack/pkg/http/httpcontext"
	mock_httpcontex "github.com/supermetrolog/goblack/tests/mocks/pkg/http/interfaces/httpcontext"
)

func TestParam(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpCtx := httpcontext.New(httptest.NewRequest("GET", "/path", nil), mock_httpcontex.NewMockResponseWriter(ctrl), map[string]string{"id": "12", "test": "1234"})
	assert.Equal(t, "12", httpCtx.Param("id"))
	assert.Equal(t, "1234", httpCtx.Param("test"))
	assert.Equal(t, "", httpCtx.Param("notExistKey"))

}
