package models

import (
	"fmt"
	"os"
	"testing"

	// "github.com/MiftahSalam/gin-blog/users/services"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/joho/godotenv"
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

func TestMain(m *testing.M) {
	common.LogI.Println("Test main users start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}
	db := common.Init()
	Init(db)

	AuthoMigrate()

	exitVal := m.Run()

	CleanUpAfterTest()

	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	os.Exit(exitVal)

	common.LogI.Println("Test main users end")
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

func TestGetUsers(t *testing.T) {
	asserts := assert.New(t)

	users, err := GetUsers()
	asserts.NoError(err, "should get all users")
	asserts.Equal(UserMockNumber, len(users))
}

func TestFindOneUser(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range UsersMock {
		userActual, err := FindOneUser(user.ID)
		asserts.NoError(err, "%v should exist", user.Username)
		asserts.Equal(user, userActual, "user should equal")
	}
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

func TestSaveOneUser(t *testing.T) {
	asserts := assert.New(t)

	image := fmt.Sprintf("http://image/%v.jpg", UserMockNumber+1)
	user := UserModel{
		Username: fmt.Sprintf("user%v", UserMockNumber+1),
		Email:    fmt.Sprintf("user%v@linkedin.com", UserMockNumber+1),
		Bio:      fmt.Sprintf("bio%v", UserMockNumber+1),
		Image:    &image,
	}
	user.SetPassword("123456")

	err := SaveOne(&user)
	asserts.NoError(err, "user %v should created", user)

	userSaved, errFind := FindOneUser(user)
	asserts.NoError(errFind, "user %v should exist", user)
	asserts.Equal(user, userSaved, "user %v should equal", user)
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
	asserts.True(UsersMock[0].IsFollowing(UsersMock[1]), "%v should be follow %v", UsersMock[0].Username, UsersMock[1].Username)

	UsersMock[0].Following(UsersMock[2])
	asserts.Equal(2, len(UsersMock[0].GetFollowing()), "user0 following users len should equal 2 ")
	asserts.True(UsersMock[0].IsFollowing(UsersMock[2]), "%v should be follow %v", UsersMock[0].Username, UsersMock[2].Username)

	follwedUser1, follwedUser2 := UsersMock[0].GetFollowing()[0], UsersMock[0].GetFollowing()[1]
	asserts.Equal(UsersMock[1], follwedUser1, "%v should same with %v", UsersMock[1], follwedUser1)
	asserts.Equal(UsersMock[2], follwedUser2, "%v should same with %v", UsersMock[2], follwedUser2)

	UsersMock[0].UnFollow(UsersMock[1])
	asserts.Equal(1, len(UsersMock[0].GetFollowing()), "user0 following users len should equal 1 ")
	asserts.False(UsersMock[0].IsFollowing(UsersMock[1]), "%v should not follow %v", UsersMock[0].Username, UsersMock[1].Username)
}

func TestDeleteOneUser(t *testing.T) {
	asserts := assert.New(t)

	var userToDelete = UsersMock[UserMockNumber-int(CurrentRecordCount)-1]
	err := DeleteOneUsers(UserModel{ID: userToDelete.ID})
	asserts.NoError(err, "should success deleted one user: %v", userToDelete.Username)
	_, err = FindOneUser(userToDelete)
	asserts.Error(err, "user %v should not found", userToDelete.Username)
}
