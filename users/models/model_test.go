package models

import (
	"fmt"
	"testing"

	// "github.com/MiftahSalam/gin-blog/users/services"

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

	err := userModelMock.CheckPassword("")
	asserts.Error(err, "empty password should return error")

	err = userModelMock.SetPassword("")
	asserts.Error(err, "empty password cannot be set null")

	err = userModelMock.SetPassword("123456")
	asserts.NoError(err, "password should be set successfully")
	asserts.NotEqual(userModelMock.PasswordHash, "123456", "hashed password should not equal to password")

	err = userModelMock.CheckPassword("12345")
	asserts.Error(err, "password should be checked and not valid")

	err = userModelMock.CheckPassword("123456")
	asserts.NoError(err, "password should be checked and valid")
}

func TestUserUpdate(t *testing.T) {
	asserts := assert.New(t)

	user := UsersMock[0]
	image := fmt.Sprintf("http://image/%vupdated.jpg", 0)
	userUpdate := UserModel{
		Email: fmt.Sprintf("user%vupdated@linkedin.com", 0),
		Bio:   fmt.Sprintf("bio%vupdated", 0),
		Image: &image,
	}
	userUpdate.SetPassword("654321")

	err := user.Update(&userUpdate)
	asserts.NoError(err, "user %v should updated", user)

	var userUpdated UserModel
	errFind := db.Where(user).First(&userUpdated).Error
	asserts.NoError(errFind, "user %v should exist", user)
	asserts.Equal(user, userUpdated, "user %v should equal", user)
}

func TestFollowing(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range UsersMock {
		t.Run(fmt.Sprintf("Test Get Following user: %v", user.Username), func(t *testing.T) {
			followingUser := user.GetFollowing()
			asserts.Empty(followingUser, "following user should be empty")
		})
	}

	UsersMock[0].Following(UsersMock[1])
	asserts.Equal(1, len(UsersMock[0].GetFollowing()), "user0 following users len should equal 1 ")
	asserts.True(UsersMock[0].IsFollowing(UsersMock[1]), "%v should be follow %v", UsersMock[0].Username, UsersMock[0].Username)
}
