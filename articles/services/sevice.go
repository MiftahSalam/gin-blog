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
	"gorm.io/gorm"
)

func ArticleCreate(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.LogI.Println("recover from panic", err)
			c.JSON(http.StatusInternalServerError, common.NewError("server", errors.New("oopss something went wrong")))
		}
	}()

	_, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, common.NewError("access", errors.New("user not login")))
		return
	}

	common.LogI.Println("req body", c.Request.Body)

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

func ArticleUpdate(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.LogI.Println("recover from panic", err)
			c.JSON(http.StatusInternalServerError, common.NewError("server", errors.New("oopss something went wrong")))
		}
	}()

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

	articleModel, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("article", errors.New("article not found")))
		return
	}

	articleValidator := validator.NewArticleModelValidator()
	if err := articleValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	articleValidator.ArticleModel.ID = articleModel.ID
	if err := articleModel.Update(&articleValidator.ArticleModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.ArticleSerializer{C: c, ArticleModel: articleModel}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}

func ArticleDelete(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("article", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("article", errors.New("user not login")))
		return
	}
	user := curUserModel.(UserModels.UserModel)
	userModel := ArticleModels.GetArticleUserModel(user)

	common.LogI.Println("userModel id", userModel.ID)
	err := ArticleModels.DeleteArticleModel(&ArticleModels.ArticleModel{Slug: slug, AuthorID: userModel.ID})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("article", errors.New("article not found")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": "Deleted"})
}

func ArticleFavorite(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("article", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("article", errors.New("user not login")))
		return
	}

	userModel := curUserModel.(UserModels.UserModel)
	article, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("article", errors.New("article not found")))
		return
	}

	err = article.FavoriteBy(ArticleModels.GetArticleUserModel(userModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.ArticleSerializer{C: c, ArticleModel: article}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}

func ArticleUnFavorite(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("article", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("article", errors.New("user not login")))
		return
	}

	userModel := curUserModel.(UserModels.UserModel)
	article, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("article", errors.New("article not found")))
		return
	}

	articleUserModel := ArticleModels.GetArticleUserModel(userModel)
	err = article.UnFavoriteBy(&articleUserModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.ArticleSerializer{C: c, ArticleModel: article}
	c.JSON(http.StatusOK, gin.H{"article": serializer.Response()})
}

func ArticleCommentCreate(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("comment", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("comment", errors.New("user not login")))
		return
	}

	commentValidator := validator.NewCommentModelValidator()
	if err := commentValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	article, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("article not found")))
		return
	}

	commentValidator.CommentModel.Article = article
	if err := ArticleModels.SaveOne(&commentValidator.CommentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.CommentSerializer{C: c, CommentModel: commentValidator.CommentModel}
	c.JSON(http.StatusCreated, gin.H{"comment": serializer.Response()})
}

func ArticleCommentDelete(c *gin.Context) {
	comment_id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("comment", errors.New("invalid comment id")))
		return
	}

	id := uint(comment_id)

	common.LogI.Println("comment id", id)

	err = ArticleModels.DeleteCommentModel(&ArticleModels.CommentModel{Model: gorm.Model{ID: id}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("comment not found")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": "Deleted"})
}

func ArticleCommentList(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, common.NewError("comments", errors.New("invalid slug")))
		return
	}

	curUserModel, _ := c.Get("user")
	if curUserModel == nil {
		c.JSON(http.StatusUnauthorized, common.NewError("comments", errors.New("user not login")))
		return
	}

	article, err := ArticleModels.FindOneArticle(&ArticleModels.ArticleModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("article not found")))
		return
	}
	comments, err := article.GetComments()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", err))
		return
	}

	serializer := serializers.CommentsSerializer{C: c, Comments: comments}
	c.JSON(http.StatusOK, gin.H{"comments": serializer.Response()})

}

func TagList(c *gin.Context) {
	tagsModel, err := ArticleModels.GetAllTags()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tags", err))
		return
	}

	serializer := serializers.TagsSerializer{C: c, Tags: tagsModel}
	c.JSON(http.StatusOK, gin.H{"tags": serializer.Response()})
}
