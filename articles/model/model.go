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
	db.AutoMigrate(&ArticleModel{})
	db.AutoMigrate(&CommentModel{})
	db.AutoMigrate(&FavoriteModel{})
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

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error

	return err
}

func FindOneArticle(condition interface{}) (ArticleModel, error) {
	db := common.GetDB()
	var model ArticleModel

	err := db.Preload("Author.UserModel").Preload("Tags").First(&model, condition).Error

	return model, err
}

func FindArticles(tag, author, favorited string, limit, offset int) ([]ArticleModel, int64, error) {
	db := common.GetDB()
	var articles []ArticleModel
	var count int64

	// user, err := userModel.FindOneUser(&userModel.UserModel{Username: author})
	// if err == nil {
	// 	common.LogI.Println("user", user)
	// } else {
	// 	common.LogI.Println("user err", err)
	// }
	// err = db.Preload("Tags", "tag = ?", tag).
	// 	Preload("Author.UserModel", "username = ?", author).
	// 	Joins("FavoriteModel.ArticleUserModel.UserModel", "username = ?", favorited).
	// 	Offset(offset_int).
	// 	Limit(limit_int).Find(&articles).Count(&count).Error

	tx := db.Begin()
	var tagModel TagModel

	tx.Where(TagModel{Tag: tag}).First(&tagModel)
	tx.Model(&ArticleModel{}).Where("tag IN ?", tag).Association("Tags").Find(&articles)
	err := tx.Commit().Error

	common.LogI.Println("tagModel", tagModel)
	// common.LogI.Println("articles", articles)
	// common.LogI.Println("count", count)
	// common.LogI.Println("err", err)

	return articles, count, err
}

func getAllTags() ([]TagModel, error) {
	db := common.GetDB()
	var tags []TagModel
	err := db.Find(&tags).Error

	return tags, err
}

func (article *ArticleModel) getComments() ([]CommentModel, error) {
	db := common.GetDB()
	var comments []CommentModel

	err := db.Model(&ArticleModel{}).Preload("Comments").Find(&comments, ArticleModel{
		Slug: article.Slug,
	}).Error

	return comments, err
}

func (article *ArticleModel) setTags(tags []string) error {
	db := common.GetDB()
	var tagList []TagModel

	for _, tag := range tags {
		var tagModel TagModel
		err := db.FirstOrCreate(&tagModel, TagModel{Tag: tag}).Error

		if err != nil {
			return err
		}
		tagList = append(tagList, tagModel)
	}
	article.Tags = tagList
	return nil
}

func (article *ArticleModel) favoriteCount() int64 {
	var count int64
	db := common.GetDB()

	db.Model(&FavoriteModel{}).Where(FavoriteModel{
		FavoriteID: article.ID,
	}).Count(&count)

	return count
}

func (article *ArticleModel) isFavoriteBy(user *ArticleUserModel) bool {
	db := common.GetDB()
	var favourite FavoriteModel

	db.Where(FavoriteModel{
		FavoriteByID: user.ID,
		FavoriteID:   article.ID,
	}).First(&favourite)

	return favourite.ID != 0
}

func (article *ArticleModel) favoriteBy(user ArticleUserModel) error {
	db := common.GetDB()
	var favourite FavoriteModel

	err := db.FirstOrCreate(&favourite, &FavoriteModel{
		FavoriteByID: user.ID,
		FavoriteID:   article.ID,
	}).Error

	return err
}

func (article *ArticleModel) unFavoriteBy(user *ArticleUserModel) error {
	db := common.GetDB()
	var favourite FavoriteModel

	err := db.Unscoped().Where(FavoriteModel{
		FavoriteID:   article.ID,
		FavoriteByID: user.ID,
	}).Delete(&favourite).Error

	return err
}

func (user *ArticleUserModel) getArticleFeed(limit, offset int) ([]ArticleModel, int, error) {
	db := common.GetDB()
	var articles []ArticleModel

	tx := db.Begin()
	followings := user.UserModel.GetFollowing()
	var articleUserModelIds []uint

	for _, following := range followings {
		articleUserModel := GetArticleUserModel(following)
		articleUserModelIds = append(articleUserModelIds, articleUserModel.ID)
	}
	tx.Preload("Author.UserModel").Preload("Tags").Where("author_id IN (?)", articleUserModelIds).Offset(offset).Limit(limit).Find(&articles)
	err := tx.Commit().Error

	// common.LogI.Println("followings", followings)
	// common.LogI.Println("count", count)
	// common.LogI.Println("err", err)

	return articles, len(articles), err
}
