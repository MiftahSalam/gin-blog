package services

import (
	"errors"
	"net/http"
	"strconv"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	serializers "github.com/MiftahSalam/gin-blog/articles/serializers"
	validator "github.com/MiftahSalam/gin-blog/articles/validators"
	"github.com/MiftahSalam/gin-blog/common"
	UserModels "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func ArticleCreate(c *gin.Context) {
	_, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, common.NewError("access", errors.New("user not login")))
		return
	}

	articleModelValidator := validator.NewArticleModelValidator()
	if err := articleModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	articleModelValidator.ArticleModel.Slug = slug.Make(articleModelValidator.ArticleModel.Title)
	if err := ArticleModels.SaveOne(&articleModelValidator.ArticleModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.ArticleSerializer{C: c, ArticleModel: articleModelValidator.ArticleModel}
	c.JSON(http.StatusCreated, gin.H{"article": serializer.Response()})
}

//todo Article list

func ArticleFeed(c *gin.Context) {
	limit := c.DefaultQuery("limit", "0")
	offset := c.DefaultQuery("offset", "0")

	common.LogI.Println("query limit", limit)
	common.LogI.Println("query offset", offset)

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("access", errors.New("user not login")))
		return
	}
	currentUser := curUserModel.(UserModels.UserModel)

	if currentUser.ID < 1 {
		c.JSON(http.StatusUnauthorized, common.NewError("access", errors.New("user not login")))
		return
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 0
	}
	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	articleUserModel := ArticleModels.GetArticleUserModel(currentUser)
	articles, articleCount, err := articleUserModel.GetArticleFeed(limit_int, offset_int)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", err))
	}

	serializer := serializers.ArticlesSerializer{C: c, Articles: articles}
	c.JSON(http.StatusOK, gin.H{
		"articles":      serializer.Response(),
		"articlesCount": articleCount,
	})
}

func ArticleRetrieve(c *gin.Context) {
	slug := c.Param("slug")

	// common.LogI.Println("slug", slug)

	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("article", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("article", errors.New("user not login")))
		return
	}

	if slug == "feed" {
		ArticleFeed(c)
		return
	}

	articleModel, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("article", errors.New("article not found")))
		return
	}

	serializer := serializers.ArticleSerializer{C: c, ArticleModel: articleModel}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}
