package routers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/gin-blog/common"
	userServices "github.com/MiftahSalam/gin-blog/users/services"

	articleModels "github.com/MiftahSalam/gin-blog/articles/model"
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

			// a.Equal(articleModels.ArticlesMock[0].Title, jsonResp.Articles[0].Title)
			// a.Equal(articleModels.ArticlesMock[0].Body, jsonResp.Articles[0].Body)
			// a.Equal(articleModels.ArticlesMock[0].Author.UserModel.Username, jsonResp.Articles[0].Author.Username)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (all articles limit 2): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?limit=2",
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

			a.Equal(uint(2), uint(len(jsonResp.Articles)))
			a.Equal(uint(2), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (all articles offset 2): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?offset=1",
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

			a.Equal(uint(3), uint(len(jsonResp.Articles)))
			a.Equal(uint(3), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (list by tag mock): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?tag=mock",
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

			a.Equal(uint(1), uint(len(jsonResp.Articles)))
			a.Equal(uint(1), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (list by author0): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?author=user0",
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

			a.Equal(uint(2), uint(len(jsonResp.Articles)))
			a.Equal(uint(2), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (list by author1): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?author=user1",
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

			a.Equal(uint(1), uint(len(jsonResp.Articles)))
			a.Equal(uint(1), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (list by author11): Get Article List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/article/?author=user11",
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

			a.Equal(uint(0), uint(len(jsonResp.Articles)))
			a.Equal(uint(0), jsonResp.ArticleCount)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (list by favorite user1): Get Article List Test",
			Init: func(req *http.Request) {
				articleModels.ArticlesMock[0].FavoriteBy(articleModels.ArticleUsersModelMock[1])
			},
			Url:             "/article/?favorited=user1",
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
			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(uint(1), uint(len(jsonResp.Articles)))
			a.Equal(uint(1), jsonResp.ArticleCount)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Articles[0].Author.Username)

			articleModels.ArticlesMock[0].UnFavoriteBy(&articleModels.ArticleUsersModelMock[1])

		},
	},
}
