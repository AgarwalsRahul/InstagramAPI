package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIError(t *testing.T) {
	e := NewAPIError(http.StatusBadGateway, "HIJACKED", "HIJACKED", nil)
	assert.Equal(t, http.StatusBadGateway, e.Status)
	assert.Equal(t, "HIJACKED", e.Message)
	assert.Equal(t, "HIJACKED", e.ErrorCode)
	assert.Nil(t, e.Details)

	e = NewAPIError(http.StatusBadRequest, "BAD_REQUEST", "BAD_REQUEST", Params{"name": "Invalid name provided"})
	assert.NotNil(t, e.Details)
	assert.Equal(t, http.StatusBadRequest, e.Status)
	assert.IsTypef(t, Params{}, e.Details, "Object not same type", nil)
}
