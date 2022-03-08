package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var image_url = "https://golang.org/doc/gopher/frontpage.png"
var userModelMock UserModel = UserModel{
	ID:           2,
	Username:     "miftah",
	Email:        "salam-miftah@gmail.com",
	Bio:          "Hello world",
	Image:        &image_url,
	PasswordHash: "",
}

func TestUserModel(t *testing.T) {
	asserts := assert.New(t)

	err := userModelMock.checkPassword("")
	asserts.Error(err, "empty password should return error")

	err = userModelMock.setPassword("")
	asserts.Error(err, "empty password cannot be set null")

	err = userModelMock.setPassword("123456")
	asserts.NoError(err, "password should be set successfully")
	asserts.NotEqual(userModelMock.PasswordHash, "123456", "hashed password should not equal to password")

	err = userModelMock.checkPassword("12345")
	asserts.Error(err, "password should be checked and not valid")

	err = userModelMock.checkPassword("123456")
	asserts.NoError(err, "password should be checked and valid")
}
