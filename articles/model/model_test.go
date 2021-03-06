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

	err := godotenv.Load("../../.test.env")
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
		Title:       "My Article 3",
		Slug:        slug.Make("My Article 3"),
		Description: "This article is about article 3",
		Body:        "Article 3 is en example of creating article",
		Author:      articleUser0Model,
		AuthorID:    articleUser0Model.ID,
	})

	asserts.NoError(err, "article 0 should created")
}

func TestSaveComment(t *testing.T) {
	// t.Skip()
	asserts := assert.New(t)

	common.LogI.Println("ArticlesMock[0]", ArticlesMock[0])
	common.LogI.Println("ArticleUsersModelMock", ArticleUsersModelMock[1])

	err := SaveOne(&CommentModel{
		Article:   ArticlesMock[0],
		ArticleID: ArticlesMock[0].ID,
		Author:    ArticleUsersModelMock[1],
		AuthorID:  ArticleUsersModelMock[1].ID,
		Body:      "this is comment one for article 0 from author 1",
	})

	asserts.NoError(err, "comment for article 0 should created")
}

func TestFindOneArticle(t *testing.T) {
	asserts := assert.New(t)

	article, err := FindOneArticle(&ArticleModel{
		Slug: ArticlesMock[0].Slug,
	})

	// common.LogI.Println("ArticlesMock0", ArticlesMock[0])

	if err != nil {
		common.LogE.Fatal("cannot find article with error: ", err)
		return
	}

	// common.LogI.Println("article found", article)

	asserts.Equal(ArticlesMock[0].Author.UserModel, article.Author.UserModel)
	asserts.Equal(ArticlesMock[0].Title, article.Title)
	asserts.Equal(ArticlesMock[0].Body, article.Body)

	article, err = FindOneArticle(&ArticleModel{
		Slug: ArticlesMock[1].Slug,
	})

	// common.LogI.Println("ArticlesMock0", ArticlesMock[0])

	if err != nil {
		common.LogE.Fatal("cannot find article with error: ", err)
		return
	}

	// common.LogI.Println("article found", article)

	asserts.Equal(ArticlesMock[1].Author.UserModel, article.Author.UserModel)
	asserts.Equal(ArticlesMock[1].Title, article.Title)
	asserts.Equal(ArticlesMock[1].Body, article.Body)
}

func TestFavourite(t *testing.T) {
	asserts := assert.New(t)

	//first favourite count check
	asserts.Equal(ArticlesMock[0].FavoriteCount(), int64(0))

	//first favourite by user check 1
	asserts.False(ArticlesMock[0].IsFavoriteBy(&ArticleUsersModelMock[1]))

	//assign favourite by
	err := ArticlesMock[0].FavoriteBy(ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 2
	asserts.Equal(ArticlesMock[0].FavoriteCount(), int64(1))

	//first favourite by user check 2
	asserts.True(ArticlesMock[0].IsFavoriteBy(&ArticleUsersModelMock[1]))

	//unfavourite by
	err = ArticlesMock[0].UnFavoriteBy(&ArticleUsersModelMock[1])
	asserts.NoError(err)

	//favourite count check 3
	asserts.Equal(ArticlesMock[0].FavoriteCount(), int64(0))

	//first favourite by user check 3
	asserts.False(ArticlesMock[0].IsFavoriteBy(&ArticleUsersModelMock[1]))
}

func TestFindArticles(t *testing.T) {
	asserts := assert.New(t)

	//assign favourite by
	err := ArticlesMock[0].FavoriteBy(ArticleUsersModelMock[1])
	asserts.NoError(err)

	//test find all articles
	articles, count, err := FindArticles("", "", "", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(4), count)
	asserts.Equal(4, len(articles))

	//test find all articles limit 2
	articles, count, err = FindArticles("", "", "", 2, 0)
	asserts.NoError(err)
	asserts.Equal(int64(4), count)
	asserts.Equal(2, len(articles))
	common.LogI.Println("articles len", len(articles))
	common.LogI.Println("count", count)

	//test find all articles offset 2
	articles, count, err = FindArticles("", "", "", 0, 2)
	asserts.NoError(err)
	asserts.Equal(2, len(articles))
	asserts.Equal(int64(4), count)

	//test find all articles offset 3 limit 2
	articles, count, err = FindArticles("", "", "", 2, 3)
	asserts.NoError(err)
	asserts.Equal(1, len(articles))
	asserts.Equal(int64(4), count)

	//test find articles by tag mock
	articles, count, err = FindArticles("mock", "", "", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(1), count)
	asserts.Equal(ArticlesMock[0].Title, articles[0].Title)

	//test find articles by tag new
	articles, count, err = FindArticles("new", "", "", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(1), count)
	asserts.Equal(ArticlesMock[1].Title, articles[0].Title)

	//test find articles by author user0
	articles, count, err = FindArticles("", "user0", "", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(2), count)
	asserts.Equal(ArticlesMock[0].Title, articles[0].Title)
	asserts.Equal("My Article 3", articles[1].Title)

	//test find articles by author user1
	articles, count, err = FindArticles("", "user1", "", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(1), count)
	asserts.Equal(ArticlesMock[1].Title, articles[0].Title)

	//test find articles by favorited user1
	articles, count, err = FindArticles("", "", "user1", 0, 0)
	asserts.NoError(err)
	asserts.Equal(int64(1), count)
	asserts.Equal(ArticlesMock[0].Title, articles[0].Title)
	asserts.Equal(ArticlesMock[0].Author.UserModel.Username, articles[0].Author.UserModel.Username)
}

func TestGetAllTags(t *testing.T) {
	asserts := assert.New(t)

	tags, err := GetAllTags()

	asserts.NoError(err)
	asserts.Equal(len(TagsMock), len(tags))
}

func TestGetArticleComments(t *testing.T) {
	// t.Skip()
	asserts := assert.New(t)

	comments0, err := ArticlesMock[0].GetComments()

	// if err == nil {
	// 	common.LogI.Println("article0", comments0)
	// } else {
	// 	common.LogI.Println("article0 err", err)
	// }

	asserts.NoError(err)
	asserts.Equal(1, len(comments0))
}

func TestGetArticleFeed(t *testing.T) {
	asserts := assert.New(t)
	//follow two of avail user mock
	ArticleUsersModelMock[0].UserModel.Following(ArticleUsersModelMock[1].UserModel)
	ArticleUsersModelMock[0].UserModel.Following(ArticleUsersModelMock[2].UserModel)

	//get article feeds
	articles, count, err := ArticleUsersModelMock[0].GetArticleFeed(0, 0)

	asserts.NoError(err)
	asserts.Equal(4, count)
	asserts.Equal(4, len(articles))

	articles, count, err = ArticleUsersModelMock[0].GetArticleFeed(1, 0)
	asserts.NoError(err)
	asserts.Equal(1, len(articles))
	asserts.Equal(4, count)

	articles, count, err = ArticleUsersModelMock[0].GetArticleFeed(0, 1)
	asserts.NoError(err)
	asserts.Equal(3, len(articles))
	asserts.Equal(4, count)

	articles, count, err = ArticleUsersModelMock[0].GetArticleFeed(1, 1)
	asserts.NoError(err)
	asserts.Equal(1, len(articles))
	asserts.Equal(4, count)

	articles, count, err = ArticleUsersModelMock[0].GetArticleFeed(1, 3)
	asserts.NoError(err)
	asserts.Equal(1, len(articles))
	asserts.Equal(4, count)

	articles, count, err = ArticleUsersModelMock[0].GetArticleFeed(-1, 3)
	/*
		if err == nil {
			common.LogI.Println("article feeds", articles)
			common.LogI.Println("article feeds len", len(articles))
			common.LogI.Println("count", count)
		} else {
			common.LogI.Println("article feeds err", err)
		}
	*/
	asserts.NoError(err)
	asserts.Equal(1, len(articles))
	asserts.Equal(4, count)
}

func TestUpdateArticle(t *testing.T) {
	asserts := assert.New(t)

	err := ArticlesMock[1].Update(&ArticleModel{
		Description: "This is my article with title My Article 1 after updated",
		Title:       "My Article 1 updated",
	})

	asserts.NoError(err)

	updated_article, _ := FindOneArticle(&ArticleModel{
		Title: ArticlesMock[1].Title,
	})

	// common.LogI.Println("updated_article err", err)
	// common.LogI.Println("updated_article", updated_article)

	asserts.Equal(ArticlesMock[1].Description, updated_article.Description)
	asserts.Equal("My Article 1 updated", updated_article.Title)
	asserts.Equal(slug.Make("My Article 1 updated"), updated_article.Slug)
}

func TestDeleteCommentModel(t *testing.T) {
	// t.Skip()
	asserts := assert.New(t)

	//create new comment to article 0
	err := SaveOne(&CommentModel{
		Article:   ArticlesMock[0],
		ArticleID: ArticlesMock[0].ID,
		Author:    ArticleUsersModelMock[1],
		AuthorID:  ArticleUsersModelMock[1].ID,
		Body:      "this is comment two for article 0 from author1",
	})
	asserts.NoError(err)

	articles0, err := FindOneArticle(&ArticleModel{
		Slug: ArticlesMock[0].Slug,
	})
	asserts.NoError(err)

	if err == nil {
		comments0, err := articles0.GetComments()
		asserts.NoError(err)
		if err == nil {
			asserts.Equal(2, len(comments0))
			err = DeleteCommentModel(&comments0[0])
			asserts.NoError(err)
			articles0, err = FindOneArticle(&ArticleModel{
				Slug: ArticlesMock[0].Slug,
			})
			asserts.NoError(err)
			comments0, err = articles0.GetComments()
			asserts.NoError(err)
			asserts.Equal(1, len(comments0))
		}
	}
}

func TestDeleteArticleModel(t *testing.T) {
	asserts := assert.New(t)

	err := DeleteArticleModel(&ArticleModel{
		Slug: ArticlesMock[0].Slug,
	})

	asserts.NoError(err)
}
