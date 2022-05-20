package users

import (
	"github.com/MiftahSalam/gin-blog/users/services"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup) {
	router.POST("/", services.Register)
	router.POST("/login", services.Login)
}

func UsersAuth(router *gin.RouterGroup) {
	router.GET("/", services.GetUsers)
	router.GET("/following", services.GetUsersFollowing)
	router.PUT("/", services.UpdateUser)
}

func Profile(router *gin.RouterGroup) {
	router.GET("/:username", services.GetUserProfile)
	router.POST("/:username/follow", services.FollowUser)
	router.DELETE("/:username/follow", services.UnFollowUser)
}
