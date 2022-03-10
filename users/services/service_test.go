package services

import (
	"fmt"
	"testing"

	"github.com/MiftahSalam/gin-blog/users/models"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	asserts := assert.New(t)

	users, err := GetUsers()
	asserts.NoError(err, "should get all users")
	asserts.Equal(models.UserMockNumber, len(users))
}

func TestFindOneUser(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range models.UsersMock {
		userActual, err := FindOneUser(&user)
		asserts.NoError(err, "%v should exist", user.Username)
		asserts.Equal(user, userActual, "user should equal")
	}
}

func TestSaveOneUser(t *testing.T) {
	asserts := assert.New(t)

	image := fmt.Sprintf("http://image/%v.jpg", models.UserMockNumber+1)
	user := models.UserModel{
		Username: fmt.Sprintf("user%v", models.UserMockNumber+1),
		Email:    fmt.Sprintf("user%v@linkedin.com", models.UserMockNumber+1),
		Bio:      fmt.Sprintf("bio%v", models.UserMockNumber+1),
		Image:    &image,
	}
	user.SetPassword("123456")

	err := SaveOne(&user)
	asserts.NoError(err, "user %v should created", user)

	userSaved, errFind := FindOneUser(user)
	asserts.NoError(errFind, "user %v should exist", user)
	asserts.Equal(user, userSaved, "user %v should equal", user)
}

func TestDeleteOneUser(t *testing.T) {
	asserts := assert.New(t)

	var userToDelete = models.UsersMock[models.UserMockNumber-1]
	err := DeleteOneUsers(models.UserModel{ID: userToDelete.ID})
	asserts.NoError(err, "should success deleted one user: %v", userToDelete.Username)
	_, err = FindOneUser(userToDelete)
	asserts.Error(err, "user %v should not found", userToDelete.Username)
}
