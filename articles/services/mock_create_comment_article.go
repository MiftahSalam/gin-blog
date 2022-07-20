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

var MockArticleCommentCreateTest = []MockTests{
	{
		"error unauthorized (no user loged in): ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article0"})
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[0].Body,
		}},
		http.StatusUnauthorized,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"comment":"user not login"}}`, string(response_body))
		},
	},
	{
		"error bad request (no data body): ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article0"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
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
		"error bad request (no slug): ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"comment":"invalid slug"}}`, string(response_body))
		},
	},
	{
		"error not found: ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-ad-article0"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[0].Body,
		}},
		http.StatusNotFound,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"comment":"article not found"}}`, string(response_body))
		},
	},
	{
		"no error: ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[0].Body,
		}},
		http.StatusCreated,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleCommentsMock[0].Body, jsonResp.Comment.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
		},
	},
	{
		"no error comment again: ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[1].Body,
		}},
		http.StatusCreated,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleCommentsMock[1].Body, jsonResp.Comment.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
		},
	},
	{
		"no error comment from another user: ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[0].Body,
		}},
		http.StatusCreated,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleCommentsMock[0].Body, jsonResp.Comment.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[1].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
		},
	},
	{
		"no error comment to another article: ArticleCommentCreate Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article2"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{"comment": {
			"body": ArticleCommentsMock[1].Body,
		}},
		http.StatusCreated,
		func(_ *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// common.LogI.Println("response_body", string(response_body))

			var jsonResp ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(ArticleCommentsMock[1].Body, jsonResp.Comment.Body)
			a.Equal(ArticleModels.ArticleUsersModelMock[1].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
			ArticleCommentsMock[1].ID = jsonResp.Comment.ID
		},
	},
}
