package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	// "net/url"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var MockArticleCommentDeleteTest = []MockTests{
	{
		"error (no param provided): ArticleCommentDelete Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"comment":"invalid comment id"}}`, string(response_body))
		},
	},
	{
		"error (invalid comment id): ArticleCommentDelete Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "id", Value: "sdfsdf"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"comment":"invalid comment id"}}`, string(response_body))
		},
	},
	{
		"error (comment not found): ArticleCommentDelete Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "id", Value: "90000"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusNotFound,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"comment":"comment not found"}}`, string(response_body))
		},
	},
	{
		"no error: ArticleCommentDelete Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{
				Key:   "id",
				Value: strconv.FormatInt(int64(ArticleCommentsMock[1].ID), 10),
			})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"comment":"Deleted"}`, string(response_body))
		},
	},
}
