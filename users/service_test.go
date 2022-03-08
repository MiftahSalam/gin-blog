package users

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOneUser(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range usersMock {
		userActual, err := FindOneUser(&user)
		asserts.NoError(err, "%v should exist", user.Username)
		asserts.Equal(user, userActual, "user should equal")
	}
}

func TestSaveOneUser(t *testing.T) {
	asserts := assert.New(t)

	image := fmt.Sprintf("http://image/%v.jpg", userMockNumber+1)
	user := UserModel{
		Username: fmt.Sprintf("user%v", userMockNumber+1),
		Email:    fmt.Sprintf("user%v@linkedin.com", userMockNumber+1),
		Bio:      fmt.Sprintf("bio%v", userMockNumber+1),
		Image:    &image,
	}
	user.setPassword("123456")

	err := SaveOne(&user)
	asserts.NoError(err, "user %v should created", user)

	userSaved, errFind := FindOneUser(user)
	asserts.NoError(errFind, "user %v should exist", user)
	asserts.Equal(user, userSaved, "user %v should equal", user)
}

func TestUserUpdate(t *testing.T) {
	asserts := assert.New(t)

	user := usersMock[0]
	image := fmt.Sprintf("http://image/%vupdated.jpg", 0)
	userUpdate := UserModel{
		Email: fmt.Sprintf("user%vupdated@linkedin.com", 0),
		Bio:   fmt.Sprintf("bio%vupdated", 0),
		Image: &image,
	}
	userUpdate.setPassword("654321")

	err := user.Update(&userUpdate)
	asserts.NoError(err, "user %v should updated", user)

	userUpdated, errFind := FindOneUser(user)
	asserts.NoError(errFind, "user %v should exist", user)
	asserts.Equal(user, userUpdated, "user %v should equal", user)
}

func TestFollowing(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range usersMock {
		t.Run(fmt.Sprintf("Test Get Following user: %v", user.Username), func(t *testing.T) {
			followingUser := user.GetFollowing()
			asserts.Empty(followingUser, "following user should be empty")
		})
	}

	usersMock[0].following(usersMock[1])
	asserts.Equal(1, len(usersMock[0].GetFollowing()), "user0 following users len should equal 1 ")
	asserts.True(usersMock[0].isFollowing(usersMock[1]), "%v should be follow %v", usersMock[0].Username, usersMock[0].Username)
}
