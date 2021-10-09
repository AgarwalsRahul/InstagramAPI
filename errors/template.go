package errors

type (
	// Params is used to replace placeholders in an error template with the corresponding values.
	Params map[string]interface{}
)

// NewAPIError creates a new APIError with the given HTTP status code, error code, and parameters for replacing placeholders in the error template.
// The param can be nil, indicating there is no need for placeholder replacement.
func NewAPIError(status int, code string, message string, params Params) *APIError {
	err := &APIError{
		Status:    status,
		ErrorCode: code,
		Message:   message,
		Details:   params,
	}
	return err
}
