package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	// ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var MockArticlesListTest = []MockTests{
	{
		"no error (all articles): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/")
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
			a.Equal(uint(5), jsonResp.ArticleCount)
			a.Equal(uint(5), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (all articles limit 2): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=2")
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
			a.Equal(uint(2), jsonResp.ArticleCount)
			a.Equal(uint(2), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (all articles offset 2): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?offset=2")
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
			a.Equal(uint(3), jsonResp.ArticleCount)
			a.Equal(uint(3), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (list by tag mock): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?tag=mock&limit=0&offset=0")
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
			a.Equal(uint(2), jsonResp.ArticleCount)
			a.Equal(uint(2), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (list by author 0): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?author=user0&limit=0&offset=0")
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
			a.Equal(uint(3), jsonResp.ArticleCount)
			a.Equal(uint(3), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (list by author 1): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?author=user1&limit=0&offset=0")
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
			a.Equal(uint(1), jsonResp.ArticleCount)
			a.Equal(uint(1), uint(len(jsonResp.Articles)))
		},
	},
	{
		"no error (list by favorite): ArticlesList Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?favorited=user1&limit=0&offset=0")
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
			a.Equal(uint(0), jsonResp.ArticleCount)
			a.Equal(uint(0), uint(len(jsonResp.Articles)))
		},
	},
}
