package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_NewRouter(t *testing.T) {

	hr := NewRouter([]Route{})
	assert.Nil(t, hr)

	route1 := Route{
		Name:   "User",
		Method: "GET",
		Path:   "/users/:id",
	}
	hr = NewRouter([]Route{route1})
	assert.NotNil(t, hr)
}
