package model

import (
	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      ArticleUserModel
	AuthorID    uint
	Tags        []TagModel     `gorm:"many2many:article_tags"`
	Comments    []CommentModel `gorm:"ForeignKey:ArticleID"`
}

type ArticleUserModel struct {
	gorm.Model
	UserModel      userModel.UserModel
	UserModelID    uint            `gorm:"column:user_id"`
	ArticleModels  []ArticleModel  `gorm:"ForeignKey:AuthorID"`
	FavoriteModels []FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}

type TagModel struct {
	gorm.Model
	Tag           string         `gorm:"unique_index"`
	ArticleModels []ArticleModel `gorm:"many2many:article_tags"`
}

type CommentModel struct {
	gorm.Model
	Article   ArticleModel
	ArticleID uint
	Author    ArticleUserModel
	AuthorID  uint
	Body      string `gorm:"size:2048"`
}

type FavoriteModel struct {
	gorm.Model
	Favorite     ArticleModel
	FavoriteID   uint
	FavoriteBy   ArticleUserModel
	FavoriteByID uint
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&ArticleUserModel{})
}

func GetArticleUserModel(user userModel.UserModel) ArticleUserModel {
	var articleUserModel ArticleUserModel

	if user.ID == 0 {
		return articleUserModel
	}
	db := common.GetDB()
	db.Where(&ArticleUserModel{
		UserModelID: user.ID,
	}).FirstOrCreate(&articleUserModel)
	articleUserModel.UserModel = user

	return articleUserModel
}
