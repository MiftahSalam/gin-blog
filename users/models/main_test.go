package models

import (
	"os"
	"testing"

	"github.com/MiftahSalam/gin-blog/common"

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
	Init(db)

	AuthoMigrate()

	exitVal := m.Run()

	CleanUpAfterTest()

	os.Exit(exitVal)

	common.LogI.Println("Test main users end")
}
