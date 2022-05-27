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
	ArticleModel ArticleModels.ArticleModel `json:"-"`
}

type CommentModelValidator struct {
	Comment struct {
		Body string `form:"body" json:"body" binding:"max=2048"`
	} `json:"comment"`
	CommentModel ArticleModels.CommentModel `json:"-"`
}

func NewArticleModelValidator() ArticleModelValidator {
	return ArticleModelValidator{}
}

func NewArticleModelValidatorFillWith(article ArticleModels.ArticleModel) ArticleModelValidator {
	articleModelValidator := NewArticleModelValidator()
	articleModelValidator.Article.Title = article.Title
	articleModelValidator.Article.Description = article.Description
	articleModelValidator.Article.Body = article.Body
	articleModelValidator.ArticleModel = article

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

	articleModelValidator.ArticleModel.Slug = slug.Make(articleModelValidator.ArticleModel.Title)
	articleModelValidator.ArticleModel.Title = articleModelValidator.Article.Title
	articleModelValidator.ArticleModel.Description = articleModelValidator.Article.Description
	articleModelValidator.ArticleModel.Body = articleModelValidator.Article.Body
	articleModelValidator.ArticleModel.Author = ArticleModels.GetArticleUserModel(UserModels.UserModel(currentUser))
	articleModelValidator.ArticleModel.SetTags(articleModelValidator.Article.Tags)

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

	commentModelValidator.CommentModel.Body = commentModelValidator.Comment.Body
	commentModelValidator.CommentModel.Author = ArticleModels.GetArticleUserModel(UserModels.UserModel(currentUser))

	return nil
}
