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

var MockCreateArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: Create Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:    "/article/",
			Method: "POST",
			Body: fmt.Sprintf(`{"article":{"title":"%v","description":"%v","body":"%v"}}`,
				articleServices.ArticlesMock[0].Title,
				articleServices.ArticlesMock[0].Description,
				articleServices.ArticlesMock[0].Body,
			),
			ResponseCode:    http.StatusCreated,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.ArticleResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)
			// common.LogI.Println("jsonResp", jsonResp)
			a.Equal(articleServices.ArticlesMock[0].Title, jsonResp.Article.Title)
			a.Equal(articleServices.ArticlesMock[0].Body, jsonResp.Article.Body)
			a.Equal(articleModels.ArticleUsersModelMock[0].UserModel.Username, jsonResp.Article.Author.Username)
		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (no body data provided): Create Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/",
			Method:          "POST",
			Body:            "",
			ResponseCode:    http.StatusInternalServerError,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusInternalServerError",
		},
		ResponseTest: func(_ *httptest.ResponseRecorder, _ *assert.Assertions) {
			// response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

		},
	},
	{
		UserMockTest: userServices.MockTests{
			TestName: "error (invalid body format): Create Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             "/article/",
			Method:          "POST",
			Body:            "fasdlkjrawer",
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
			TestName: "error unauthorized (no token present): Create Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Content-Type", "application/json")
			},
			Url:    "/article/",
			Method: "POST",
			Body: fmt.Sprintf(`{"article":{"title":"%v","description":"%v","body":"%v"}}`,
				articleServices.ArticlesMock[0].Title,
				articleServices.ArticlesMock[0].Description,
				articleServices.ArticlesMock[0].Body,
			),
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
}
