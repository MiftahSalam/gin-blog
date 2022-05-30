package routers

import (
	"github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/gin-gonic/gin"
)

func Articles(router *gin.RouterGroup) {
	// router.GET("/", services.ArticleList) //todo
	router.GET("/:slug", services.ArticleRetrieve)
	router.GET("/:slug/comments", services.ArticleCommentList)
}
