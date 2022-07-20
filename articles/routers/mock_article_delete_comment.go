package routers

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/gin-blog/common"
	userModels "github.com/MiftahSalam/gin-blog/users/models"
	userServices "github.com/MiftahSalam/gin-blog/users/services"

	// articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	articleServices "github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/stretchr/testify/assert"
)

var MockDeleteCommentArticle = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: DeleteCommentArticle Article Test",
			Init: func(req *http.Request) {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(userModels.UsersMock[0].ID)))
			},
			Url:             fmt.Sprintf("/article/my-article1/comment/%v", articleServices.ArticleCommentsMock[1].ID),
			Method:          "DELETE",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			a.Equal(`{"comment":"Deleted"}`, string(response_body))
		},
	},
}
