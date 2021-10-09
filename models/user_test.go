package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_OK(t *testing.T) {
	user := User{}

	err := user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Name")

	user.Name = "Rahul"
	err = user.OK()
	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "email")

	user.Email = "abc@abc.com"
	err = user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "password")

	user.Password = "password"
	err = user.OK()
	assert.Nil(t, err)
}
