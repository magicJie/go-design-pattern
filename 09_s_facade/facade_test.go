package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Login(130010101010, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.LoginOrRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Register(1300101010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}
