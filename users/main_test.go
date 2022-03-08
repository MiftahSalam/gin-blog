package users

import (
	"fmt"
	"os"
	"testing"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const userMockNumber = 3

var db *gorm.DB
var usersMock []UserModel

func createUsersMock(n int) []UserModel {
	var ret []UserModel

	if n < 2 {
		panic("user mock count should be greater or equal to 2")
	}

	for i := 0; i < n; i++ {
		image := fmt.Sprintf("http://image/%v.jpg", i)
		userModel := UserModel{
			Username: fmt.Sprintf("user%v", i),
			Email:    fmt.Sprintf("user%v@linkedin.com", i),
			Bio:      fmt.Sprintf("bio%v", i),
			Image:    &image,
		}
		userModel.setPassword("123456")
		// common.LogI.Println("create user", userModel)
		db.Create(&userModel)
		ret = append(ret, userModel)
	}

	return ret
}

func cleanUpAfterTest() {
	common.LogI.Println("clean up start")

	for _, user := range usersMock {
		common.LogI.Println("clean up user", user)

		db.Debug().Unscoped().Delete(FollowModel{}, &FollowModel{
			FollowedByID: user.ID,
		})

		db.Debug().Unscoped().Delete(FollowModel{}, &FollowModel{
			FollowingID: user.ID,
		})
	}

	db.Debug().Where("username LIKE ?", "%user%").Delete(UserModel{})
	common.LogI.Println("cleaned up")
}

func TestMain(m *testing.M) {
	common.LogI.Println("Test main users start")

	err := godotenv.Load("../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}
	db = common.Init()

	AuthoMigrate()
	usersMock = createUsersMock(userMockNumber)

	exitVal := m.Run()

	cleanUpAfterTest()
	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	os.Exit(exitVal)

	common.LogI.Println("Test main users end")
}
