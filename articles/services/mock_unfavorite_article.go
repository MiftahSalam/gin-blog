package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var MockArticleUnFavoriteTest = []MockTests{
	{
		"error (no slug provided): MockArticleUnFavoriteTest Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"invalid slug"}}`, string(response_body))
		},
	},
	{
		"error (article not found): ArticleFavorite Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "fe7ed"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusNotFound,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"article not found"}}`, string(response_body))
		},
	},

	{
		"error unauthorized (no user loged in): ArticleFavorite Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article0"})
		},
		map[string]map[string]interface{}{},
		http.StatusUnauthorized,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"user not login"}}`, string(response_body))
		},
	},
	{
		"no error : ArticleFavorite Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleModels.ArticlesMock[1].Title, jsonResp.Article.Title)
			a.Equal(ArticleModels.ArticlesMock[1].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticlesMock[1].Author.UserModel.Username, jsonResp.Article.Author.Username)
			a.Equal(uint(1), jsonResp.Article.FavoritesCount)
		},
	},
	{
		"no error (favorited by another): ArticleFavorite Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[2].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleModels.ArticlesMock[1].Title, jsonResp.Article.Title)
			a.Equal(ArticleModels.ArticlesMock[1].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticlesMock[1].Author.UserModel.Username, jsonResp.Article.Author.Username)
			a.Equal(uint(0), jsonResp.Article.FavoritesCount)
		},
	},
}
