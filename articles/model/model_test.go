package model_test

import (
	"os"
	"testing"

	articlesModel "github.com/MiftahSalam/gin-blog/articles/model"
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

	articlesModel.AutoMigrate()
	articlesModel.Init(db)

	exitVal := m.Run()

	articlesModel.CleanUpAfterTest()
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
		articleUsersModel := articlesModel.GetArticleUserModel(user)
		asserts.Equal(articleUsersModel.UserModel, user)
	}
}

func TestSaveArticle(t *testing.T) {
	asserts := assert.New(t)

	articleUser0Model := articlesModel.GetArticleUserModel(userModel.UsersMock[0])

	err := articlesModel.SaveOne(&articlesModel.ArticleModel{
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

	err := articlesModel.SaveOne(&articlesModel.CommentModel{
		Article:   articlesModel.ArticlesMock[0],
		ArticleID: articlesModel.ArticlesMock[0].ID,
		Author:    articlesModel.ArticleUsersModelMock[1],
		AuthorID:  articlesModel.ArticleUsersModelMock[1].ID,
		Body:      "this is comment for article 0",
	})

	asserts.NoError(err, "comment for article 0 should created")
}

func TestFindOneArticle(t *testing.T) {
	asserts := assert.New(t)

	article, err := articlesModel.FindOneArticle(&articlesModel.ArticleModel{
		Slug: articlesModel.ArticlesMock[0].Slug,
	})

	common.LogI.Println("ArticlesMock0", articlesModel.ArticlesMock[0])

	if err != nil {
		common.LogE.Fatal("cannot find article with error: ", err)
		return
	}

	common.LogI.Println("article found", article)

	asserts.Equal(articlesModel.ArticlesMock[0].Author.UserModel, article.Author.UserModel)
	asserts.Equal(articlesModel.ArticlesMock[0].Title, article.Title)
	asserts.Equal(articlesModel.ArticlesMock[0].Body, article.Body)
}

func TestFavourite(t *testing.T) {
	asserts := assert.New(t)

	//first favourite count check
	asserts.Equal(articlesModel.ArticlesMock[0].FavoriteCount(), int64(0))

	//first favourite by user check 1
	asserts.False(articlesModel.ArticlesMock[0].IsFavoriteBy(&articlesModel.ArticleUsersModelMock[1]))

	//assign favourite by
	err := articlesModel.ArticlesMock[0].FavoriteBy(articlesModel.ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 2
	asserts.Equal(articlesModel.ArticlesMock[0].FavoriteCount(), int64(1))

	//first favourite by user check 2
	asserts.True(articlesModel.ArticlesMock[0].IsFavoriteBy(&articlesModel.ArticleUsersModelMock[1]))

	//unfavourite by
	err = articlesModel.ArticlesMock[0].UnFavoriteBy(&articlesModel.ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 3
	asserts.Equal(articlesModel.ArticlesMock[0].FavoriteCount(), int64(0))

	//first favourite by user check 3
	asserts.False(articlesModel.ArticlesMock[0].IsFavoriteBy(&articlesModel.ArticleUsersModelMock[1]))
}
