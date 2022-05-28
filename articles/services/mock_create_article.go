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

var MockArticleCreateTest = []MockTests{
	{
		"no error: ArticleCreate Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[0].Title,
			"description": ArticlesMock[0].Description,
			"body":        ArticlesMock[0].Body,
			"tagList":     ArticleModels.TagsMock,
		}},
		http.StatusCreated,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))
			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticlesMock[0].Title, jsonResp.Article.Title)
			a.Equal(ArticlesMock[0].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	{
		"error unauthorized (no user loged in): ArticleCreate Test",
		func(c *gin.Context) {
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[1].Title,
			"description": ArticlesMock[1].Description,
			"body":        ArticlesMock[1].Body,
			"tagList":     ArticleModels.TagsMock,
		}},
		http.StatusUnauthorized,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			common.LogI.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"access":"user not login"}}`, string(response_body))

		},
	},
	{
		"error bad request (no data body): ArticleCreate Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			common.LogI.Println("response_body", string(response_body))
			// a.Equal(`{"errors":{"access":"user not login"}}`, string(response_body))
		},
	},
	{
		"no error althought some data body fields not provided: ArticleCreate Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"article": {
			"title": ArticlesMock[1].Title,
			"body":  ArticlesMock[1].Body,
		}},
		http.StatusCreated,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))
			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticlesMock[1].Title, jsonResp.Article.Title)
			a.Equal(ArticlesMock[1].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
}
