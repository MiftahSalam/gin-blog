package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/gin-blog/common"
	userServices "github.com/MiftahSalam/gin-blog/users/services"

	// articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	articleServices "github.com/MiftahSalam/gin-blog/articles/services"
	"github.com/stretchr/testify/assert"
)

var MockGetTagList = []RouterMockTest{
	{
		UserMockTest: userServices.MockTests{
			TestName: "no error: Get Tag List Test",
			Init: func(req *http.Request) {
			},
			Url:             "/tags/",
			Method:          "GET",
			Body:            "",
			ResponseCode:    http.StatusOK,
			ResponsePattern: "",
			Msg:             "valid data and should return StatusOK",
		},
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// common.LogI.Println("response_body", string(response_body))

			var jsonResp articleServices.TagsResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// common.LogI.Println("jsonResp", jsonResp)

		},
	},
}
