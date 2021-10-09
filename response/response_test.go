package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_GenericResponse(t *testing.T) {

	resp := GenericResponse(http.StatusCreated, "User created successfully", nil)
	assert.Equal(t, resp.Data, nil)
	assert.Equal(t, resp.StatusCode(), http.StatusCreated)
	assert.Equal(t, resp.Status, http.StatusCreated)
}

func TestResponse_WriteResponse(t *testing.T) {
	recorder := httptest.NewRecorder()

	// This should make the app panic
	assert.Panics(t, func() { WriteResponse(recorder, nil) }, "Panic failed")

	// This should pass
	resp := GenericResponse(200, "User created successfully.", map[string]string{"a": "xyz"})
	assert.NotPanics(t, func() { WriteResponse(recorder, resp) })
}
