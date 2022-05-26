package validators

import (
	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	UserModels "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ArticleModelValidator struct {
	Article struct {
		Title       string   `form:"title" json:"title" binding:"required,min=4"`
		Description string   `form:"description" json:"description" binding:"max=2048"`
		Body        string   `form:"body" json:"body"`
		Tags        []string `form:"tagList" json:"tagList"`
	} `json:"article"`
	articleModel ArticleModels.ArticleModel `json:"-"`
}

type CommentModelValidator struct {
	Comment struct {
		Body string `form:"body" json:"body" binding:"max=2048"`
	} `json:"comment"`
	commentModel ArticleModels.CommentModel `json:"-"`
}

func NewArticleModelValidator() ArticleModelValidator {
	return ArticleModelValidator{}
}

func NewArticleModelValidatorFillWith(article ArticleModels.ArticleModel) ArticleModelValidator {
	articleModelValidator := NewArticleModelValidator()
	articleModelValidator.Article.Title = article.Title
	articleModelValidator.Article.Description = article.Description
	articleModelValidator.Article.Body = article.Body
	articleModelValidator.articleModel = article

	for _, tag := range article.Tags {
		articleModelValidator.Article.Tags = append(articleModelValidator.Article.Tags, tag.Tag)
	}

	return articleModelValidator
}

func (articleModelValidator *ArticleModelValidator) Bind(c *gin.Context) error {
	currentUser := c.MustGet("user").(UserModels.UserModel)

	err := common.Bind(c, articleModelValidator)
	if err != nil {
		return err
	}

	articleModelValidator.articleModel.Slug = slug.Make(articleModelValidator.articleModel.Title)
	articleModelValidator.articleModel.Title = articleModelValidator.Article.Title
	articleModelValidator.articleModel.Description = articleModelValidator.Article.Description
	articleModelValidator.articleModel.Body = articleModelValidator.Article.Body
	articleModelValidator.articleModel.Author = ArticleModels.GetArticleUserModel(currentUser)
	articleModelValidator.articleModel.SetTags(articleModelValidator.Article.Tags)

	return nil
}

func NewCommentModelValidator() CommentModelValidator {
	return CommentModelValidator{}
}

func (commentModelValidator *CommentModelValidator) Bind(c *gin.Context) error {
	currentUser := c.MustGet("user").(UserModels.UserModel)

	err := common.Bind(c, commentModelValidator)
	if err != nil {
		return err
	}

	commentModelValidator.commentModel.Body = commentModelValidator.Comment.Body
	commentModelValidator.commentModel.Author = ArticleModels.GetArticleUserModel(currentUser)

	return nil
}
