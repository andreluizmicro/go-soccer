package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const name = "John Doe"
const email = "john@gmail.com"

func TestNewuser(t *testing.T) {
	user, err := NewUser(name, email, "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser(name, email, "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
