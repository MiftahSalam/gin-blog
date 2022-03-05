package main

import (
	// "log"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		common.LogE.Println("Error loading file .env")
	}
	common.LogI.Println("Test log format")

	users.CheckDotEnv()
}
