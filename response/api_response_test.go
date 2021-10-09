package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIResponse_ResponseMessage(t *testing.T) {

	r := APIResponse{
		Message: "abc",
	}
	assert.Equal(t, "abc", r.ResponseMessage())
}

func TestAPIResponse_StatusCode(t *testing.T) {

	r := APIResponse{
		Status: 200,
	}
	assert.Equal(t, 200, r.StatusCode())
}
