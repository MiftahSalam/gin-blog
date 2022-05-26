package serializers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/gothinkster/golang-gin-realworld-example-app/users"
	"github.com/stretchr/testify/assert"
)

var articleModel ArticleModels.ArticleModel

func TestMain(m *testing.M) {
	articleModel = ArticleModels.ArticleModel{
		Title:       "Create Article",
		Description: "Create Article Description",
		Body:        "Create Article Body",
		Tags: []ArticleModels.TagModel{
			{Tag: "tag1"},
			{Tag: "tag2"},
		},
	}
	articleModel.Slug = slug.Make(articleModel.Title)

	exitVal := m.Run()
	os.Exit(exitVal)
}

//hard to test because must mock ArticleUserModel
func TestArticleSerializerResponse(t *testing.T) {
	asserts := assert.New(t)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Set("user", users.UserModel{
		ID:       1,
		Username: "user1",
		Email:    "user1@gmail.com",
		Bio:      "user1 bio",
	})

	articleSerializer := ArticleSerializer{C: c, ArticleModel: articleModel}
	asserts.Equal(articleSerializer.Response().Body, articleModel.Body)
}
