package routers

import (
	"github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/gin-gonic/gin"
)

func Articles(router *gin.RouterGroup) {
	router.POST("/", services.ArticleCreate)
	router.GET("/feed", services.ArticleFeed)
	router.PUT("/:slug", services.ArticleUpdate)
	router.POST("/:slug/favorite", services.ArticleFavorite)
	router.DELETE("/:slug/favorite", services.ArticleUnFavorite)
	router.DELETE("/:slug", services.ArticleDelete)
	router.POST("/:slug/comment", services.ArticleCommentCreate)
	router.DELETE("/:slug/comment/:id", services.ArticleCommentDelete)
}

func ArticlesAnonymous(router *gin.RouterGroup) {
	router.GET("/", services.ArticleList)
	router.GET("/:slug", services.ArticleRetrieve)
	router.GET("/:slug/comments", services.ArticleCommentList)
}

func Tags(router *gin.RouterGroup) {
	router.GET("/", services.TagList)
}
