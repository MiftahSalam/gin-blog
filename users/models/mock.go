package models

import (
	"fmt"

	"github.com/MiftahSalam/gin-blog/common"
	"gorm.io/gorm"
)

var UserMockNumber = 3
var CurrentRecordCount int64

var db *gorm.DB
var UsersMock []UserModel

func Init(database *gorm.DB) {
	db = database

	UsersMock = createUsersMock(UserMockNumber)

}

func createUsersMock(n int) []UserModel {
	var ret []UserModel

	if n < 2 {
		panic("user mock count should be greater or equal to 2")
	}

	//count existing record
	var users []UserModel
	db.Find(&users).Count(&CurrentRecordCount)
	common.LogI.Println("CurrentRecordCount", CurrentRecordCount)
	UserMockNumber += int(CurrentRecordCount)

	for i := 0; i < n; i++ {
		image := fmt.Sprintf("http://image/%v.jpg", i)
		userModel := UserModel{
			Username: fmt.Sprintf("user%v", i),
			Email:    fmt.Sprintf("user%v@linkedin.com", i),
			Bio:      fmt.Sprintf("bio%v", i),
			Image:    &image,
		}
		userModel.SetPassword("123456")
		// common.LogI.Println("create user", userModel)
		db.Create(&userModel)
		ret = append(ret, userModel)
	}

	return ret
}

func CleanUpAfterTest() {
	common.LogI.Println("clean up users start")

	for _, user := range UsersMock {
		// common.LogI.Println("clean up user", user)

		db.Unscoped().Delete(FollowModel{}, &FollowModel{
			FollowedByID: user.ID,
		})

		db.Unscoped().Delete(FollowModel{}, &FollowModel{
			FollowingID: user.ID,
		})
	}

	db.Debug().Where("username LIKE ?", "%user%").Delete(UserModel{})
	common.LogI.Println("clean up users end")
}
