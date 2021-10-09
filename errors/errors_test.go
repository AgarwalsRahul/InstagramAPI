package errors

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {

	e := InternalServerError(errors.New(""))
	assert.Equal(t, http.StatusInternalServerError, e.Status)
}

func TestNotFound(t *testing.T) {

	e := NotFound("/xyz")
	assert.Equal(t, http.StatusNotFound, e.Status)
}

func TestUnauthorized(t *testing.T) {

	e := Unauthorized("Invalid email or password provided")
	assert.Equal(t, http.StatusUnauthorized, e.Status)
}

func TestGenericError(t *testing.T) {

	e := GenericError(450, Params{"x": "xyz"}, "INVALID_TOKEN", "Provide a valid user token")
	assert.Equal(t, 450, e.Status)
	assert.Equal(t, "INVALID_TOKEN", e.ErrorCode)
}

func TestResponse_WriteErrorResponse(t *testing.T) {
	recorder := httptest.NewRecorder()

	// This should make the app panic
	assert.Panics(t, func() { WriteErrorResponse(recorder, nil) }, "Panic failed")

	// This should pass
	apiErr := GenericError(http.StatusBadRequest, Params{"a": "xyz"}, "BAD_REQUEST", "Please provide a valid response")
	assert.NotPanics(t, func() { WriteErrorResponse(recorder, apiErr) })
}
