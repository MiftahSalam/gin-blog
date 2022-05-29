package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	// "net/url"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var MockArticleDeleteTest = []MockTests{
	{
		"error (no slug provided): ArticlesDelete Test",
		func(c *gin.Context) {
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"invalid slug"}}`, string(response_body))
		},
	},

	{
		"error (article not found): ArticlesDelete Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "fe7ed"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[0].UserModel)
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
		"error unauthorized (no user loged in): ArticlesDelete Test",
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
	{
		"no error : ArticlesDelete Test",
		func(c *gin.Context) {
			c.Params = append(c.Params, gin.Param{Key: "slug", Value: "my-article1"})
			c.Set("user", ArticleModels.ArticleUsersModelMock[1].UserModel)
		},
		map[string]map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"article":"Deleted"}`, string(response_body))
		},
	},
}
