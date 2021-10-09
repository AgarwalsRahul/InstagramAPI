package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosts_OK(t *testing.T) {
	post := Post{}

	err := post.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Caption")

	post.Caption = "Caption"
	err = post.OK()
	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "ImageUrl")

	post.ImageUrl = "https://www.gogle.com"
	err = post.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "UserId")

	post.UserId = "asdajsdjsad456ccac"
	err = post.OK()
	assert.Nil(t, err)
}
