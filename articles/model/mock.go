package model

import (
	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"gorm.io/gorm"
)

var db *gorm.DB
var ArticleUsersModelMock []ArticleUserModel
var CurrentArticleUsersRecordCount int64
var ArticleUsersMockNumber int64

// var ArticlesMock []ArticleModel
// var ArticlesMockNumber = 3
// var CurrentArticleRecordCount int64

func Init(database *gorm.DB) {
	db = database
	userModel.Init(db)

	ArticleUsersModelMock = createArticleUsersModelMock(len(userModel.UsersMock))
	// ArticlesMock = createArticlesMock(ArticlesMockNumber)

}

func createArticleUsersModelMock(n int) []ArticleUserModel {
	var ret []ArticleUserModel

	if n < 2 {
		panic("article mock count should be greater or equal to 2")
	}

	//count existing record
	var articleUsersModel []ArticleUserModel
	db.Find(&articleUsersModel).Count(&CurrentArticleUsersRecordCount)
	common.LogI.Println("CurrentArticleUsersRecordCount", CurrentArticleUsersRecordCount)
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
		common.LogI.Println("created article user", articleUser)
		ret = append(ret, articleUser)
	}

	return ret
}

// func createArticlesMock(n int) []ArticleModel {
// 	var ret []ArticleModel

// 	if n < 2 {
// 		panic("article mock count should be greater or equal to 2")
// 	}

// 	//count existing record
// 	var articles []ArticleModel
// 	db.Find(&articles).Count(&CurrentArticleRecordCount)
// 	common.LogI.Println("CurrentArticleRecordCount", CurrentArticleRecordCount)
// 	ArticlesMockNumber += int(CurrentArticleRecordCount)

// 	for i := 0; i < n; i++ {
// 		article := ArticleModel{
// 			Title: fmt.Sprintf("My article%v", i),
// 			Description:      fmt.Sprintf("This is my article with title My Article %v", i),
// 			Body:  fmt.Sprintf("Article %v is created with gin gonic with title My Article %v", i, i),
// 		}
// 		article.Slug = slug.Make(article.Title)
// 		common.LogI.Println("created article", article)
// 		db.Create(&article)
// 		ret = append(ret, article)
// 	}

// 	return ret
// }

func CleanUpAfterTest() {
	common.LogI.Println("clean up article start")

	userModel.CleanUpAfterTest()

	for _, user := range userModel.UsersMock {
		common.LogI.Println("clean up article user", user)

		db.Debug().Where("user_id = ?", user.ID).Delete(ArticleUserModel{})
	}

	common.LogI.Println("cleaned up article")
}
