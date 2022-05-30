package routers

import (
	"github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/gin-gonic/gin"
)

func Articles(router *gin.RouterGroup) {
	// router.GET("/", services.ArticleList) //todo
	router.POST("/", services.ArticleCreate)
	router.GET("/:slug", services.ArticleRetrieve)
	router.PUT("/:slug", services.ArticleUpdate)
	router.GET("/:slug/comments", services.ArticleCommentList)
}
