package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	ArticleSerializers "github.com/MiftahSalam/gin-blog/articles/serializers"
	"github.com/MiftahSalam/gin-blog/common"
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

type ArticleCommentsResponse struct {
	Comments []struct {
		ArticleSerializers.CommentResponse
	} `json:"comments"`
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
		Body: "this is comment for article",
	},
	{
		Body: "this is comment for article second",
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

func MockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbyte, err := json.Marshal(content)
	if err != nil {
		common.LogE.Println("Cannot marshal json content")
		panic(err)
	}
	common.LogI.Println("content", string(jsonbyte))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbyte))
}
