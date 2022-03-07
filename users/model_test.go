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
}
