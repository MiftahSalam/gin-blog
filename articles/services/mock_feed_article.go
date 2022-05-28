package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var MockArticlesFeedTest = []MockTests{
	{
		"no error: ArticlesFeed Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=0&offset=0")
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)

			ArticleModels.ArticleUsersModelMock[1].UserModel.Following(ArticleModels.ArticleUsersModelMock[0].UserModel)
			ArticleModels.ArticleUsersModelMock[1].UserModel.Following(ArticleModels.ArticleUsersModelMock[2].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticlesResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(uint(4), jsonResp.ArticleCount)
			a.Equal(uint(4), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error limit 1 offset 0: ArticlesFeed Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=1&offset=0")
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticlesResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			common.LogI.Println("jsonResp", jsonResp)
			a.Equal(uint(1), jsonResp.ArticleCount)
			a.Equal(uint(1), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error limit 2 offset 0: ArticlesFeed Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=2&offset=0")
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticlesResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			common.LogI.Println("jsonResp", jsonResp)
			a.Equal(uint(2), jsonResp.ArticleCount)
			a.Equal(uint(2), uint(len(jsonResp.Articles)))
		},
	},
	{
		"error unauthorized (no user loged in): ArticlesFeed Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=0&offset=0")
		},
		map[string]map[string]interface{}{},
		http.StatusUnauthorized,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			common.LogI.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"access":"user not login"}}`, string(response_body))

		},
	},
}
