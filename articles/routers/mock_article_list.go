package routers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/gin-blog/common"
	// userModels "github.com/MiftahSalam/gin-blog/users/models"
	userServices "github.com/MiftahSalam/gin-blog/users/services"

	// articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	articleServices "github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/stretchr/testify/assert"
)

var MockListArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/",
			Method:          "GET",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticlesResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(uint(4), uint(len(jsonResp.Articles)))
			a.Equal(uint(4), jsonResp.ArticleCount)

			// a.Equal(articleModels.ArticlesMock[0].Title, jsonResp.Article.Title)
			// a.Equal(articleModels.ArticlesMock[0].Body, jsonResp.Article.Body)
			// a.Equal(articleModels.ArticlesMock[0].Author.UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	/*
		{
			UserMockTest: userServices.MockTests{
				TestName: "no error (another article): Get Article Test",
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
				a.Equal(articleModels.ArticlesMock[1].Title, jsonResp.Article.Title)
				a.Equal(articleModels.ArticlesMock[1].Body, jsonResp.Article.Body)
				a.Equal(articleModels.ArticlesMock[1].Author.UserModel.Username, jsonResp.Article.Author.Username)
			},
		},
		{
			UserMockTest: userServices.MockTests{
				TestName: "no error (articles feed): Get Article Test",
				Init: func(req *http.Request) {
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[1].ID)))

					articleModels.ArticleUsersModelMock[1].UserModel.Following(articleModels.ArticleUsersModelMock[0].UserModel)
					articleModels.ArticleUsersModelMock[1].UserModel.Following(articleModels.ArticleUsersModelMock[2].UserModel)

				},
				Url:             "/article/feed",
				Method:          "GET",
				Body:            "",
				ResponseCode:    http.StatusOK,
				ResponsePattern: "",
				Msg:             "valid data and should return StatusOK",
			},
			ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
				response_body, _ := ioutil.ReadAll(w.Body)

				common.LogI.Println("response_body", string(response_body))

				var jsonResp articleServices.ArticlesResponse
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
			UserMockTest: userServices.MockTests{
				TestName: "error (article not found): Get Article Test",
				Init: func(req *http.Request) {
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[1].ID)))
				},
				Url:             "/article/sdsd",
				Method:          "GET",
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
				TestName: "error unauthorized (no token present): Get Article Test",
				Init: func(req *http.Request) {
				},
				Url:             "/article/my-article0",
				Method:          "GET",
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
	*/
}
