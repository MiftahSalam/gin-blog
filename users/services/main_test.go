package services

import (
	"os"
	"testing"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	common.LogI.Println("Test main users start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}
	db := common.Init()
	models.Init(db)

	models.AuthoMigrate()

	exitVal := m.Run()

	models.CleanUpAfterTest()

	os.Exit(exitVal)

	common.LogI.Println("Test main users end")
}
