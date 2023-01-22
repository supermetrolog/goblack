package httpcontext_test

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_goblack "github.com/supermetrolog/goblack/mocks"
	"github.com/supermetrolog/goblack/pkg/http/httpcontext"
)

func TestParam(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpCtx := httpcontext.New(httptest.NewRequest("GET", "/path", nil), mock_goblack.NewMockWriter(ctrl), map[string]string{"id": "12", "test": "1234"})
	assert.Equal(t, "12", httpCtx.Param("id"))
	assert.Equal(t, "1234", httpCtx.Param("test"))
	assert.Equal(t, "", httpCtx.Param("notExistKey"))

}
