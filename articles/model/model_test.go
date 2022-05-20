package model_test

import (
	"os"
	"testing"

	articleUserModel "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	common.LogI.Println("Test Main Article start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}

	db := common.Init()

	articleUserModel.AutoMigrate()
	articleUserModel.Init(db)

	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	exitVal := m.Run()

	os.Exit(exitVal)

	common.LogI.Println("Test Main Article end")
}

func TestGetArticleUserModel(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range userModel.UsersMock {
		articleUsersModel := articleUserModel.GetArticleUserModel(user)
		asserts.Equal(articleUsersModel.UserModel.Username, user.Username)
	}
}
