package validators

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	UserModels "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
)

// func MockJSONPost(c *gin.Context, content interface{}) {
// 	c.Request.Method = "POST"
// 	c.Request.Header.Set("Content-Type", "application/json")

// 	jsonbyte, err := json.Marshal(content)
// 	if err != nil {
// 		common.LogE.Println("Cannot marshal json content")
// 		panic(err)
// 	}

// 	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbyte))
// }

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

// articleModel.Slug = slug.Make("Create Article")

func TestNewArticleModelValidator(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal(ArticleModelValidator{}, NewArticleModelValidator())
}

func TestNewArticleModelValidatorFillWith(t *testing.T) {
	asserts := assert.New(t)

	articleModelValidator := NewArticleModelValidatorFillWith(articleModel)

	asserts.Equal(articleModel, articleModelValidator.ArticleModel)

	asserts.Equal(articleModel.Title, articleModelValidator.Article.Title)
	asserts.Equal(articleModel.Description, articleModelValidator.Article.Description)
	asserts.Equal(articleModel.Body, articleModelValidator.Article.Body)
	asserts.Equal([]string{"tag1", "tag2"}, articleModelValidator.Article.Tags)
}

//hard to test because must mock ArticleUserModel
func TestArticleModelValidatorBind(t *testing.T) {
	t.Skip()
	asserts := assert.New(t)
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Set("user", UserModels.UserModel{
		ID:       1,
		Username: "user1",
		Email:    "user1@gmail.com",
		Bio:      "user1 bio",
	})

	articleModelValidator := NewArticleModelValidatorFillWith(articleModel)

	err := articleModelValidator.Bind(c)
	asserts.NoError(err)
}

func TestNewCommentModelValidator(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal(CommentModelValidator{}, NewCommentModelValidator())
}

//hard to test because must mock ArticleUserModel
func TestCommentModelValidatorBind(t *testing.T) {
	t.Skip()
	asserts := assert.New(t)
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Set("user", UserModels.UserModel{
		ID:       1,
		Username: "user1",
		Email:    "user1@gmail.com",
		Bio:      "user1 bio",
	})

	commentModelValidator := NewCommentModelValidator()

	err := commentModelValidator.Bind(c)
	asserts.NoError(err)
}
