package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "john@email", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "john@email", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John", "john@email", "123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123"))
	assert.False(t, user.ValidatePassword("1234"))
	assert.NotEqual(t, "123", user.Password)
}
