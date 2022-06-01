package model

import (
	"errors"

	"github.com/MiftahSalam/gin-blog/common"
	userModel "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gosimple/slug"
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
	Comments    []CommentModel `gorm:"ForeignKey:ID;constraint:OnDelete:CASCADE"`
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
	ArticleModels []ArticleModel `gorm:"many2many:article_tags;constraint:OnDelete:CASCADE"`
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

func DeleteArticleModel(condition interface{}) error {
	db := common.GetDB()

	article, err := FindOneArticle(condition)
	if err != nil {
		common.LogE.Println("cannot find article: ", err)
		return err
	}

	// common.LogI.Println("clean up article tags", article.Tags)
	err = db.Unscoped().Model(&article).Association("Tags").Clear()
	if err != nil {
		common.LogE.Println("cannot delete article tags: ", err)
		return err
	}

	// common.LogI.Println("clean up article favourite", article.Slug)
	err = db.Unscoped().Where(&FavoriteModel{
		FavoriteID: article.ID,
	}).Delete(FavoriteModel{}).Error
	if err != nil {
		common.LogE.Println("cannot delete article favourite: ", err)
		return err
	}

	// common.LogI.Println("clean up article comments", article.Comments)
	err = db.Unscoped().Where("article_id = ?", article.ID).Delete(CommentModel{}).Error
	if err != nil {
		common.LogE.Println("cannot delete comment ", err)
		return err
	}

	err = db.Unscoped().Where(condition).Delete(&ArticleModel{}).Error

	return err
}

func DeleteCommentModel(condition interface{}) error {
	db := common.GetDB()
	tx := db.Debug().Unscoped().Where(condition).Delete(&CommentModel{})
	err := tx.Error
	row_affected := tx.RowsAffected

	if row_affected == 0 {
		err = errors.New("comment not found")
	}
	common.LogI.Println("row_affected", row_affected)

	return err
}

func FindArticles(tag, author, favorited string, limit, offset int) ([]ArticleModel, int64, error) {
	db := common.GetDB()
	var articles []ArticleModel
	var count int64

	tx := db.Begin()
	if tag != "" {
		var tagModel TagModel
		tx.Debug().Where(TagModel{Tag: tag}).First(&tagModel)

		common.LogI.Println("tagModel.ID", tagModel.ID)

		if tagModel.ID != 0 {
			tx.Model(&tagModel).
				Preload("Tags").
				Preload("Author.UserModel").
				Offset(offset).Limit(limit).
				Association("ArticleModels").
				Find(&articles)
			count = tx.Model(&tagModel).Association("ArticleModels").Count()
		}
	} else if author != "" {
		var user userModel.UserModel
		tx.Where(userModel.UserModel{Username: author}).First(&user)
		articleUser := GetArticleUserModel(user)

		common.LogI.Println("articleUser.ID", articleUser.ID)

		if articleUser.ID != 0 {
			tx.Model(&articleUser).
				Preload("Tags").
				Preload("Author.UserModel").
				Offset(offset).Limit(limit).
				Association("ArticleModels").
				Find(&articles)
			count = tx.Model(&articleUser).Association("ArticleModels").Count()
		}
	} else if favorited != "" {
		var user userModel.UserModel
		tx.Where(userModel.UserModel{Username: favorited}).First(&user)
		articleUser := GetArticleUserModel(user)

		if articleUser.ID != 0 {
			var favoritedModel []FavoriteModel
			tx.Where(FavoriteModel{FavoriteByID: articleUser.ID}).
				Offset(offset).Limit(limit).Find(&favoritedModel)

			count = tx.Model(&articleUser).Association("FavoriteModels").Count()
			for _, favorite := range favoritedModel {
				var article ArticleModel
				tx.Model(&favorite).
					Preload("Tags").
					Preload("Author.UserModel").
					Association("Favorite").Find(&article)

				articles = append(articles, article)
			}
		}
	} else {
		db.Preload("Tags").Preload("Author.UserModel").Offset(offset).Limit(limit).Find(&articles)
		count = int64(len(articles))
	}

	err := tx.Commit().Error

	// common.LogI.Println("articles", articles)
	// common.LogI.Println("count", count)
	// common.LogI.Println("err", err)

	return articles, count, err
}

func GetAllTags() ([]TagModel, error) {
	db := common.GetDB()
	var tags []TagModel
	err := db.Find(&tags).Error

	return tags, err
}

func (article *ArticleModel) Update(data interface{}) error {
	db := common.GetDB()
	var buf_data *ArticleModel = data.(*ArticleModel)

	buf_data.Slug = slug.Make(buf_data.Title)
	err := db.Model(article).Updates(data).Error

	return err
}

func (article *ArticleModel) GetComments() ([]CommentModel, error) {
	db := common.GetDB()
	var comments []CommentModel

	err := db.Preload("Author.UserModel").Model(&CommentModel{}).Find(&comments, CommentModel{
		ArticleID: article.ID,
	}).Error

	return comments, err
}

func (article *ArticleModel) SetTags(tags []string) error {
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
	article.Update(&ArticleModel{Tags: tagList})

	return nil
}

func (article *ArticleModel) FavoriteCount() int64 {
	var count int64
	db := common.GetDB()

	db.Model(&FavoriteModel{}).Where(FavoriteModel{
		FavoriteID: article.ID,
	}).Count(&count)

	return count
}

func (article *ArticleModel) IsFavoriteBy(user *ArticleUserModel) bool {
	db := common.GetDB()
	var favourite FavoriteModel

	db.Where(FavoriteModel{
		FavoriteByID: user.ID,
		FavoriteID:   article.ID,
	}).First(&favourite)

	return favourite.ID != 0
}

func (article *ArticleModel) FavoriteBy(user ArticleUserModel) error {
	db := common.GetDB()
	var favourite FavoriteModel

	err := db.FirstOrCreate(&favourite, &FavoriteModel{
		FavoriteByID: user.ID,
		FavoriteID:   article.ID,
	}).Error

	return err
}

func (article *ArticleModel) UnFavoriteBy(user *ArticleUserModel) error {
	db := common.GetDB()
	var favourite FavoriteModel

	err := db.Unscoped().Where(FavoriteModel{
		FavoriteID:   article.ID,
		FavoriteByID: user.ID,
	}).Delete(&favourite).Error

	return err
}

func (user *ArticleUserModel) GetArticleFeed(limit, offset int) ([]ArticleModel, int, error) {
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
