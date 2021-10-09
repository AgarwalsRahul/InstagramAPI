package response

import (
	"encoding/json"
	"net/http"
)

// GenericResponse :- Means to construct a response with generic data
func GenericResponse(statusCode int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}

// WriteResponse :- Writes the response as a standard JSON response with status in APIResonse
func WriteResponse(w http.ResponseWriter, resp *APIResponse) {
	if resp == nil {
		panic("Please provide a valid response")
	}

	responseJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(resp.Status)
	w.Write(responseJSON)
}
