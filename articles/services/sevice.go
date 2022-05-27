package services

import (
	"errors"
	"net/http"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	serializer "github.com/MiftahSalam/gin-blog/articles/serializers"
	validator "github.com/MiftahSalam/gin-blog/articles/validators"
	"github.com/MiftahSalam/gin-blog/common"
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

	serializer := serializer.ArticleSerializer{C: c, ArticleModel: articleModelValidator.ArticleModel}
	c.JSON(http.StatusCreated, gin.H{"article": serializer.Response()})
}
