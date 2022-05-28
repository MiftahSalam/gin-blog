package services

import (
	"net/http/httptest"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	ArticleSerializers "github.com/MiftahSalam/gin-blog/articles/serializers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ArticleResponse struct {
	Article struct {
		ArticleSerializers.ArticleResponse
	} `json:"article"`
}

type ArticlesResponse struct {
	Articles []struct {
		ArticleSerializers.ArticleResponse
	} `json:"articles"`
	ArticleCount uint `json:"articlesCount"`
}

type ArticleCommentResponse struct {
	Comment struct {
		ArticleSerializers.CommentResponse
	} `json:"comment"`
}

var ArticlesMock = []ArticleModels.ArticleModel{
	{
		Title:       "My Article From Service",
		Description: "This is my article with title My Article From Service",
		Body:        "Article From Service is created with gin gonic with title My Article From Service",
	},
	{
		Title:       "My Article1 From Service",
		Description: "This is my article with title My Article From Service1",
		Body:        "Article From Service1 is created with gin gonic with title My Article From Service1",
	},
	{
		Title:       "My ArticleUpdated From Service1",
		Description: "This is article with title My ArticleUpdated From Service1",
		Body:        "ArticleArticleUpdated From Service1 is created with gin gonic with title My ArticleUpdated From Service1",
	},
}

var ArticleCommentsMock = []ArticleModels.CommentModel{
	{
		Body: "This is comment for article",
	},
}

var tagsMockUpdate = []string{"service", "tag"}

type MockTests struct {
	TestName     string
	Init         func(c *gin.Context)
	Data         interface{}
	ResponseCode int
	ResponseTest func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions)
}
