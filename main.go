package main

import (
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users"
	userMiddlewares "github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Migrate() {
	models.AuthoMigrate()
}
func main() {
	err := godotenv.Load()
	if err != nil {
		common.LogE.Println("Error loading file .env")
	}

	db := common.Init()
	Migrate()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			common.LogE.Fatal("get db instance error:  ", err)
		}
		sqlDB.Close()
	}()

	router := gin.Default()
	v1 := router.Group("/api")

	// v1.Use(userMiddlewares.AuthMiddleware(true))
	users.Users(v1.Group("/users"))
	v1.Use(userMiddlewares.AuthMiddleware(true))
	users.UsersAuth(v1.Group("/users"))
	users.Profile(v1.Group("/profile"))

	router.Run()
}
