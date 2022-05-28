package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	ArticleSerializers "github.com/MiftahSalam/gin-blog/articles/serializers"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ArticleResponse struct {
	Article struct {
		ArticleSerializers.ArticleResponse
	} `json:"article"`
}

type ArticlesResponse struct {
	Articles []struct {
		ArticleSerializers.ArticleResponse
	} `json:"articles"`
	ArticleCount uint `json:"articlesCount"`
}

var ArticlesMock = []ArticleModels.ArticleModel{
	{
		Title:       "My Article From Service",
		Description: "This is my article with title My Article From Service",
		Body:        "Article From Service is created with gin gonic with title My Article From Service",
	},
	{
		Title:       "My Article1 From Service",
		Description: "This is my article with title My Article From Service1",
		Body:        "Article From Service1 is created with gin gonic with title My Article From Service1",
	},
	{
		Title:       "My ArticleUpdated From Service1",
		Description: "This is article with title My ArticleUpdated From Service1",
		Body:        "ArticleArticleUpdated From Service1 is created with gin gonic with title My ArticleUpdated From Service1",
	},
}

var tagsMockUpdate = []string{"service", "tag"}

type MockTests struct {
	TestName     string
	Init         func(c *gin.Context)
	Data         interface{}
	ResponseCode int
	ResponseTest func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions)
}

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

var MockArticleRetrieveTest = []MockTests{
	{
		"no error (get one article): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article0"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
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
			a.Equal(ArticleModels.ArticlesMock[0].Title, jsonResp.Article.Title)
			a.Equal(ArticleModels.ArticlesMock[0].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticlesMock[0].Author.UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	{
		"no error (get another one article): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
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
			a.Equal(ArticleModels.ArticlesMock[1].Title, jsonResp.Article.Title)
			a.Equal(ArticleModels.ArticlesMock[1].Body, jsonResp.Article.Body)
			a.Equal(ArticleModels.ArticlesMock[1].Author.UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	{
		"no error (get another article feed): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/?limit=0&offset=0")
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "feed"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
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
		"no error (no query data provided): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Request.URL, _ = url.Parse("/")
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "feed"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
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
		"error (no slug provided): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			common.LogI.Println("jsonResp", jsonResp)
			a.Equal(`{"errors":{"article":"invalid slug"}}`, string(response_body))
		},
	},
	{
		"error (article not found): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "fe7ed"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusNotFound,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"article not found"}}`, string(response_body))
		},
	},
	{
		"error unauthorized (no user loged in): ArticlesRetrieve Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article0"})
		},
		map[string]map[string]interface{}{},
		http.StatusUnauthorized,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"user not login"}}`, string(response_body))
		},
	},
}

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
			"tagList":     tagsMockUpdate,
		}},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			common.LogI.Println("jsonResp", jsonResp)

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
			"tagList":     tagsMockUpdate,
		}},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

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
			"tagList":     tagsMockUpdate,
		}},
		http.StatusNotFound,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

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
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

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
			"tagList":     tagsMockUpdate,
		}},
		http.StatusUnauthorized,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"user not login"}}`, string(response_body))
		},
	},
}
