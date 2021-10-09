package response

// APIResponse represents a response struct for building response data
type APIResponse struct {
	// Status represents the HTTP status code
	Status int `json:"-"`
	// Message is the readable reason behind the response
	Message string `json:"message"`
	// Data specifies the additional information in the response
	Data interface{} `json:"data,omitempty"`
}

// ResponseMessage returns the API response message
func (r APIResponse) ResponseMessage() string {
	return r.Message
}

// StatusCode returns the status code of the response
func (r APIResponse) StatusCode() int {
	return r.Status
}
