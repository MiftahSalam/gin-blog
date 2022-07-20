package model

import (
	"fmt"

	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

var db *gorm.DB

var ArticleUsersModelMock []ArticleUserModel
var CurrentArticleUsersRecordCount int64
var ArticleUsersMockNumber int64

var ArticlesMock []ArticleModel
var ArticlesMockNumber = 3
var CurrentArticleRecordCount int64

var TagsMock = []string{"mock", "article", "new", "tag"}
var TagsMockNumber = 3
var CurrentTagRecordCount int64

func createArticlesMock(n int) []ArticleModel {
	var ret []ArticleModel

	if n < 2 {
		panic("article mock count should be greater or equal to 2")
	}

	//count existing record
	var articles []ArticleModel
	db.Find(&articles).Count(&CurrentArticleRecordCount)
	common.LogI.Println("CurrentArticleRecordCount", CurrentArticleRecordCount)
	ArticlesMockNumber += int(CurrentArticleRecordCount)

	for i := 0; i < n; i++ {
		article := ArticleModel{
			Title:       fmt.Sprintf("My article%v", i),
			Description: fmt.Sprintf("This is my article with title My Article %v", i),
			Body:        fmt.Sprintf("Article %v is created with gin gonic with title My Article %v", i, i),
			Author:      ArticleUsersModelMock[i],
			AuthorID:    ArticleUsersModelMock[i].ID,
		}
		article.Slug = slug.Make(article.Title)

		err := db.Create(&article).Error
		if err != nil {
			common.LogE.Fatal("failed to create article", err)
			continue
		}

		// common.LogI.Println("created article", article)

		ret = append(ret, article)
	}
	ret[0].SetTags(TagsMock[:2])
	// ret[1].SetTags(TagsMock)
	ret[1].SetTags(TagsMock[2:])

	return ret
}

func createArticleUsersModelMock(n int) []ArticleUserModel {
	var ret []ArticleUserModel

	if n < 2 {
		panic("article mock count should be greater or equal to 2")
	}

	//count existing record
	var articleUsersModel []ArticleUserModel
	db.Find(&articleUsersModel).Count(&CurrentArticleUsersRecordCount)
	// common.LogI.Println("CurrentArticleUsersRecordCount", CurrentArticleUsersRecordCount)
	ArticleUsersMockNumber += CurrentArticleUsersRecordCount

	// common.LogI.Println("UsersMock len", len(userModel.UsersMock))
	// common.LogI.Println("ArticleUserModel len", n)

	for i := 0; i < n; i++ {
		articleUser := ArticleUserModel{
			UserModel:   userModel.UsersMock[i],
			UserModelID: userModel.UsersMock[i].ID,
		}

		err := db.Create(&articleUser).Error
		if err != nil {
			common.LogE.Fatal("failed to create article user", err)
			continue
		}
		// common.LogI.Println("created article user", articleUser)
		ret = append(ret, articleUser)
	}

	return ret
}

func CleanUpAfterTest() {
	common.LogI.Println("clean up article start")

	for _, article := range ArticlesMock {
		// common.LogI.Println("clean up article tags", article.Tags)
		err := db.Unscoped().Model(&article).Association("Tags").Clear()
		if err != nil {
			common.LogE.Println("cannot delete article tags: ", err)
		}

		// common.LogI.Println("clean up article favourite", article.Slug)
		err = db.Unscoped().Where(&FavoriteModel{
			FavoriteID: article.ID,
		}).Delete(FavoriteModel{}).Error
		if err != nil {
			common.LogE.Println("cannot delete article favourite: ", err)
		}
	}
	for _, tag := range TagsMock {
		// common.LogI.Println("clean up tag", tag)

		var tagModel TagModel
		err := db.Unscoped().Delete(&tagModel, TagModel{
			Tag: tag,
		}).Error
		if err != nil {
			common.LogE.Println("cannot delete tag: ", err)
		}
	}
	err := db.Unscoped().Where("body LIKE ?", "%this is comment %").Delete(CommentModel{}).Error
	if err != nil {
		common.LogE.Println("cannot delete comment ", err)
	}
	err = db.Unscoped().Where("slug LIKE ?", "%my-article%").Delete(ArticleModel{}).Error
	if err != nil {
		common.LogE.Println("cannot delete articles ", err)
	}
	for _, user := range userModel.UsersMock {
		// common.LogI.Println("clean up article user", user)

		err = db.Unscoped().Delete(ArticleUserModel{}, &ArticleUserModel{
			UserModelID: user.ID,
		}).Error
		if err != nil {
			common.LogE.Println("cannot delete article user ", err)
		}
	}

	userModel.CleanUpAfterTest()

	common.LogI.Println("clean up article end")
}

func Init(database *gorm.DB) {
	db = database
	userModel.Init(db)

	ArticleUsersModelMock = createArticleUsersModelMock(len(userModel.UsersMock))
	ArticlesMock = createArticlesMock(len(userModel.UsersMock))

}
