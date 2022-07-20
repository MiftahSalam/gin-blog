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

var MockCreateCommentArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (article not found): Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:    "/article/my-article10/comment",
			Method: "POST",
			Body: fmt.Sprintf(`{"comment":{"body":"%v"}}`,
				articleServices.ArticleCommentsMock[0].Body,
			),
			ResponseCode:    http.StatusNotFound,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusNotFound",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"comment":"article not found"}}`, string(response_body))
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (invalid body format): Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/my-article10/comment",
			Method:          "POST",
			Body:            "sdfgsfg",
			ResponseCode:    http.StatusBadRequest,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusBadRequest",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Contains(string(response_body), "json error")
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (no data body provided): Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/my-article10/comment",
			Method:          "POST",
			Body:            "",
			ResponseCode:    http.StatusInternalServerError,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusInternalServerError",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Contains(string(response_body), "oopss")
		},
	},

	{
		UserMockTest: userServices.MockTests{
			TestName: "error unauthorized (no token present): Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
			},
			Url:             "/article/my-article10/comment",
			Method:          "POST",
			Body:            "",
			ResponseCode:    http.StatusUnauthorized,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusUnauthorized",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"error":"no token present in request"}`, string(response_body))
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:    "/article/my-article1/comment",
			Method: "POST",
			Body: fmt.Sprintf(`{"comment":{"body":"%v"}}`,
				articleServices.ArticleCommentsMock[0].Body,
			),
			ResponseCode:    http.StatusCreated,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(articleServices.ArticleCommentsMock[0].Body, jsonResp.Comment.Body)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
			articleServices.ArticleCommentsMock[0].ID = jsonResp.Comment.ID
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (create another for other article): Create Comment Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:    "/article/my-article2/comment",
			Method: "POST",
			Body: fmt.Sprintf(`{"comment":{"body":"%v"}}`,
				articleServices.ArticleCommentsMock[1].Body,
			),
			ResponseCode:    http.StatusCreated,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleCommentResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(articleServices.ArticleCommentsMock[1].Body, jsonResp.Comment.Body)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
			articleServices.ArticleCommentsMock[1].ID = jsonResp.Comment.ID
		},
	},
}
