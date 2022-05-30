package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/gin-blog/common"
	userModels "github.com/MiftahSalam/gin-blog/users/models"
	userServices "github.com/MiftahSalam/gin-blog/users/services"

	articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	articleServices "github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/stretchr/testify/assert"
)

var MockUnFavoriteArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (article not found): UnFavorite Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/my-article10/favorite",
			Method:          "DELETE",
			Body:            "",
			ResponseCode:    http.StatusNotFound,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusNotFound",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"article":"article not found"}}`, string(response_body))
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "error unauthorized (no token present): UnFavorite Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
			},
			Url:             "/article/my-artile1/favorite",
			Method:          "DELETE",
			Body:            "",
			ResponseCode:    http.StatusUnauthorized,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusUnauthorized",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"error":"no token present in request"}`, string(response_body))
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: UnFavorite Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/my-article1/favorite",
			Method:          "DELETE",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			common.LogI.Println("jsonResp", jsonResp)

			a.Equal(articleModels.ArticlesMock[1].Title, jsonResp.Article.Title)
			a.Equal(articleModels.ArticlesMock[1].Body, jsonResp.Article.Body)
			a.Equal(articleModels.ArticlesMock[1].Author.UserModel.Username, jsonResp.Article.Author.Username)
			a.Equal(uint(0), jsonResp.Article.FavoritesCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (get article favorite count): UnFavorite Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[userModels.UserMockNumber-int(userModels.CurrentRecordCount)-1].ID)))
			},
			Url:             "/article/my-article1",
			Method:          "GET",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			common.LogI.Println("jsonResp", jsonResp)
			a.Equal(uint(0), jsonResp.Article.FavoritesCount)
		},
	},
}
