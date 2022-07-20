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

var MockArticleUpdate = []MockTests{
	{
		"no error: ArticleUpdate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: ArticleModels.ArticlesMock[0].Slug})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[2].Title,
			"description": ArticlesMock[2].Description,
			"body":        ArticlesMock[2].Body,
			"tagList":     TagsMockUpdate,
		}},
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

			a.Equal(ArticlesMock[2].Title, jsonResp.Article.Title)
			a.Equal(ArticlesMock[2].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	{
		"error bad request (no slug provided): ArticleUpdate Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[2].Title,
			"description": ArticlesMock[2].Description,
			"body":        ArticlesMock[2].Body,
			"tagList":     TagsMockUpdate,
		}},
		http.StatusBadRequest,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"invalid slug"}}`, string(response_body))
		},
	},
	{
		"error not found: ArticleUpdate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "ArticleModels.ArticlesMock[0].Slug"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[2].Title,
			"description": ArticlesMock[2].Description,
			"body":        ArticlesMock[2].Body,
			"tagList":     TagsMockUpdate,
		}},
		http.StatusNotFound,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"article not found"}}`, string(response_body))
		},
	},
	{
		"error bad request (no data body): ArticleUpdate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: ArticleModels.ArticlesMock[1].Slug})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Contains(string(response_body), "key: required")
		},
	},
	{
		"error unauthorized (no user loged in): ArticleUpdate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: ArticleModels.ArticlesMock[2].Slug})
		},
		map[string]map[string]interface{}{"article": {
			"title":       ArticlesMock[2].Title,
			"description": ArticlesMock[2].Description,
			"body":        ArticlesMock[2].Body,
			"tagList":     TagsMockUpdate,
		}},
		http.StatusUnauthorized,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"user not login"}}`, string(response_body))
		},
	},
}
