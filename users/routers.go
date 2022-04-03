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
}
