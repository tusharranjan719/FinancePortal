package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	var username, id = "john doe", 15

	token, err := MakeToken(uint64(id), username)
	assert.NoError(t, err)

	if assert.NotEqual(t, token, "") {
		t.Log("succeed to create token for the user")
	}
}
