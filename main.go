package main

import (
	"log"

	"github.com/MiftahSalam/gin-blog/users"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading file .env")
	}

	users.CheckDotEnv()
}
