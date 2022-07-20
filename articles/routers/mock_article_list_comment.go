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

var MockCommentListArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: ArticleCommentList Test",
			Init: func(req *http.Request) {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[userModels.UserMockNumber-int(userModels.CurrentRecordCount)-1].ID)))
			},
			Url:             "/article/my-article1/comments",
			Method:          "GET",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleCommentsResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(uint(1), uint(len(jsonResp.Comments)))
			a.Equal(articleServices.ArticleCommentsMock[0].Body, jsonResp.Comments[0].Body)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comments[0].Author.Username)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (create new comment article1): Comment List Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[2].ID)))
			},
			Url:    "/article/my-article1/comment",
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
			a.Equal(articleModels.ArticleUsersModelMock[2].UserModel.Username, jsonResp.Comment.CommentResponse.Author.Username)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error (check again for article1): ArticleCommentList Test",
			Init: func(req *http.Request) {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[userModels.UserMockNumber-int(userModels.CurrentRecordCount)-1].ID)))
			},
			Url:             "/article/my-article1/comments",
			Method:          "GET",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleCommentsResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// common.LogI.Println("jsonResp", jsonResp)

			a.Equal(uint(2), uint(len(jsonResp.Comments)))
			a.Equal(articleServices.ArticleCommentsMock[0].Body, jsonResp.Comments[0].Body)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Comments[0].Author.Username)
			a.Equal(articleServices.ArticleCommentsMock[1].Body, jsonResp.Comments[1].Body)
			a.Equal(articleModels.ArticleUsersModelMock[2].UserModel.Username, jsonResp.Comments[1].Author.Username)
		},
	},
}
