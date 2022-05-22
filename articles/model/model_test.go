package model_test

import (
	"os"
	"testing"

	articleUserModel "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gosimple/slug"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	common.LogI.Println("Test Main Article start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}

	db := common.Init()

	articleUserModel.AutoMigrate()
	articleUserModel.Init(db)

	exitVal := m.Run()

	articleUserModel.CleanUpAfterTest()
	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	os.Exit(exitVal)

	common.LogI.Println("Test Main Article end")
}

func TestGetArticleUserModel(t *testing.T) {
	asserts := assert.New(t)

	for _, user := range userModel.UsersMock {
		articleUsersModel := articleUserModel.GetArticleUserModel(user)
		asserts.Equal(articleUsersModel.UserModel, user)
	}
}

func TestSaveArticle(t *testing.T) {
	asserts := assert.New(t)

	articleUser0Model := articleUserModel.GetArticleUserModel(userModel.UsersMock[0])

	err := articleUserModel.SaveOne(&articleUserModel.ArticleModel{
		Title:       "My Article 0",
		Slug:        slug.Make("My Article 0"),
		Description: "This article is about article 0",
		Body:        "Article 0 is en example of creating article",
		Author:      articleUser0Model,
		AuthorID:    articleUser0Model.ID,
	})

	asserts.NoError(err, "article 0 should created")
}

func TestSaveComment(t *testing.T) {
	asserts := assert.New(t)

	err := articleUserModel.SaveOne(&articleUserModel.CommentModel{
		Article:   articleUserModel.ArticlesMock[0],
		ArticleID: articleUserModel.ArticlesMock[0].ID,
		Author:    articleUserModel.ArticleUsersModelMock[1],
		AuthorID:  articleUserModel.ArticleUsersModelMock[1].ID,
		Body:      "this is comment for article 0",
	})

	asserts.NoError(err, "comment for article 0 should created")
}

func TestFindOneArticle(t *testing.T) {
	article, err := articleUserModel.FindOneArticle(&articleUserModel.ArticleModel{
		Slug: articleUserModel.ArticlesMock[0].Slug,
	})

	common.LogI.Println("ArticlesMock0 slug", articleUserModel.ArticlesMock[0].Slug)

	if err != nil {
		common.LogE.Fatal("cannot find article with error: ", err)
		return
	}

	common.LogI.Println("article found", article)

}
