package main

import (
	articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	articleRouter "github.com/MiftahSalam/gin-blog/articles/routers"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users"
	userMiddlewares "github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Migrate() {
	models.AuthoMigrate()
	articleModels.AutoMigrate()
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
	router.SetTrustedProxies(nil)

	// router.Use(func(ctx *gin.Context) {
	// 	common.LogI.Println("request", ctx.Request)
	// })
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Origin", "content-type", "Authorization"}
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/api")

	// v1.Use(userMiddlewares.AuthMiddleware(true)) //global middleware
	users.Users(v1.Group("/users"))
	users.UsersAuth(v1.Group("/user", userMiddlewares.AuthMiddleware(true)))
	users.Profile(v1.Group("/profile", userMiddlewares.AuthMiddleware(true)))

	articleRouter.Articles(v1.Group("/article", userMiddlewares.AuthMiddleware(true)))
	articleRouter.ArticlesAnonymous(v1.Group("/article", userMiddlewares.AuthMiddleware(false)))
	articleRouter.Tags(v1.Group("/tags", userMiddlewares.AuthMiddleware(false)))

	router.Run()
}
