package errors

import (
	"encoding/json"
	"net/http"
)

// InternalServerError creates a new API error representing an internal server error (HTTP 500)
func InternalServerError(err error) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusInternalServerError, defaultMsg, defaultMsg, Params{"error": err.Error()})
}

// NotFound creates a new API error representing a resource-not-found error (HTTP 404)
func NotFound(resource string) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusNotFound, defaultMsg, defaultMsg, Params{"error": resource})
}

// Unauthorized creates a new API error representing an authentication failure (HTTP 401)
func Unauthorized(err string) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusUnauthorized, defaultMsg, defaultMsg, Params{"error": err})
}

// GenericError :- For creating generic errors
func GenericError(statusCode int, values Params, messageCode, message string) *APIError {
	err := NewAPIError(statusCode, messageCode, message, nil)
	err.Details = values
	return err
}

// WriteErrorResponse :- Writes the response as a standard JSON response with status in APIError
func WriteErrorResponse(w http.ResponseWriter, apiErr *APIError) {
	if apiErr == nil {
		panic("Please provide a valid error")
	}

	errorJSON, _ := json.Marshal(apiErr)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(apiErr.Status)
	w.Write(errorJSON)
}
