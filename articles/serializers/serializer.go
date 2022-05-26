package serializers

import (
	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	UserModels "github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
	"github.com/gothinkster/golang-gin-realworld-example-app/users"
)

type ArticleSerializer struct {
	C *gin.Context
	ArticleModels.ArticleModel
}

type ArticlesSerializer struct {
	C        *gin.Context
	Articles []ArticleModels.ArticleModel
}

type ArticleUserSerializer struct {
	C *gin.Context
	ArticleModels.ArticleUserModel
}

type CommentsSerializer struct {
	C        *gin.Context
	Comments []ArticleModels.CommentModel
}

type CommentSerializer struct {
	C *gin.Context
	ArticleModels.CommentModel
}

type TagSerializer struct {
	C *gin.Context
	ArticleModels.TagModel
}

type TagsSerializer struct {
	C    *gin.Context
	Tags []ArticleModels.TagModel
}

type ArticleResponse struct {
	ID             uint                  `json:"-"`
	Title          string                `json:"title"`
	Slug           string                `json:"slug"`
	Description    string                `json:"description"`
	Body           string                `json:"body"`
	CreatedAt      string                `json:"createdAt"`
	UpdatedAt      string                `json:"updatedAt"`
	Author         users.ProfileResponse `json:"author"`
	Tags           []string              `json:"tagList"`
	Favorite       bool                  `json:"favorited"`
	FavoritesCount uint                  `json:"favoritesCount"`
}

type CommentResponse struct {
	ID        uint                  `json:"id"`
	Body      string                `json:"body"`
	CreatedAt string                `json:"createdAt"`
	UpdatedAt string                `json:"updatedAt"`
	Author    users.ProfileResponse `json:"author"`
}

func (tag *TagSerializer) Response() string {
	return tag.TagModel.Tag
}

func (tags *TagsSerializer) Response() []string {
	response := []string{}
	for _, tag := range tags.Tags {
		serializer := TagSerializer{tags.C, tag}
		response = append(response, serializer.Response())
	}

	return response
}

func (article *ArticleSerializer) Response() ArticleResponse {
	currentUser := article.C.MustGet("user").(users.UserModel)
	authorSerializer := ArticleUserSerializer{C: article.C, ArticleUserModel: article.Author}
	userArticle := ArticleModels.GetArticleUserModel(UserModels.UserModel(currentUser))
	response := ArticleResponse{
		ID:             article.ID,
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		CreatedAt:      article.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt:      article.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Author:         authorSerializer.Response(),
		Favorite:       article.IsFavoriteBy(&userArticle),
		FavoritesCount: uint(article.FavoriteCount()),
	}
	response.Tags = make([]string, 0)
	for _, tag := range article.Tags {
		serializer := TagSerializer{article.C, tag}
		response.Tags = append(response.Tags, serializer.Response())
	}

	return response
}

func (articleUser *ArticleUserSerializer) Response() users.ProfileResponse {
	response := users.ProfileSerializer{
		C:         articleUser.C,
		UserModel: users.UserModel(articleUser.ArticleUserModel.UserModel),
	}
	return response.Response()
}

func (articles *ArticlesSerializer) Response() []ArticleResponse {
	response := []ArticleResponse{}
	for _, article := range articles.Articles {
		serializer := ArticleSerializer{articles.C, article}
		response = append(response, serializer.Response())
	}

	return response
}

func (comment *CommentSerializer) Response() CommentResponse {
	authorSerializer := ArticleUserSerializer{comment.C, comment.Author}
	response := CommentResponse{
		ID:        comment.ID,
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: comment.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Author:    authorSerializer.Response(),
	}

	return response
}

func (comments *CommentsSerializer) Response() []CommentResponse {
	response := []CommentResponse{}
	for _, comment := range comments.Comments {
		serializer := CommentSerializer{comments.C, comment}
		response = append(response, serializer.Response())
	}

	return response
}
