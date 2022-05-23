package model

import (
	"os"
	"testing"

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

	AutoMigrate()
	Init(db)

	exitVal := m.Run()

	CleanUpAfterTest()
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
		articleUsersModel := GetArticleUserModel(user)
		asserts.Equal(articleUsersModel.UserModel, user)
	}
}

func TestSaveArticle(t *testing.T) {
	asserts := assert.New(t)

	articleUser0Model := GetArticleUserModel(userModel.UsersMock[0])

	err := SaveOne(&ArticleModel{
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

	err := SaveOne(&CommentModel{
		Article:   ArticlesMock[0],
		ArticleID: ArticlesMock[0].ID,
		Author:    ArticleUsersModelMock[1],
		AuthorID:  ArticleUsersModelMock[1].ID,
		Body:      "this is comment for article 0",
	})

	asserts.NoError(err, "comment for article 0 should created")
}

func TestFindOneArticle(t *testing.T) {
	asserts := assert.New(t)

	article, err := FindOneArticle(&ArticleModel{
		Slug: ArticlesMock[0].Slug,
	})

	common.LogI.Println("ArticlesMock0", ArticlesMock[0])

	if err != nil {
		common.LogE.Fatal("cannot find article with error: ", err)
		return
	}

	common.LogI.Println("article found", article)

	asserts.Equal(ArticlesMock[0].Author.UserModel, article.Author.UserModel)
	asserts.Equal(ArticlesMock[0].Title, article.Title)
	asserts.Equal(ArticlesMock[0].Body, article.Body)
}

func TestFavourite(t *testing.T) {
	asserts := assert.New(t)

	//first favourite count check
	asserts.Equal(ArticlesMock[0].favoriteCount(), int64(0))

	//first favourite by user check 1
	asserts.False(ArticlesMock[0].isFavoriteBy(&ArticleUsersModelMock[1]))

	//assign favourite by
	err := ArticlesMock[0].favoriteBy(ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 2
	asserts.Equal(ArticlesMock[0].favoriteCount(), int64(1))

	//first favourite by user check 2
	asserts.True(ArticlesMock[0].isFavoriteBy(&ArticleUsersModelMock[1]))

	//unfavourite by
	err = ArticlesMock[0].unFavoriteBy(&ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 3
	asserts.Equal(ArticlesMock[0].favoriteCount(), int64(0))

	//first favourite by user check 3
	asserts.False(ArticlesMock[0].isFavoriteBy(&ArticleUsersModelMock[1]))
}
