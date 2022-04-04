package services_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users"
	"github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/MiftahSalam/gin-blog/users/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	common.LogI.Println("Test main users services start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Error: ", err)
		panic("Cannot load env file")
	}

	db := common.Init()
	models.Init(db)
	models.AuthoMigrate()

	router = gin.New()
	router.Use(middlewares.AuthMiddleware(false))
	users.Users(router.Group("/users"))
	router.Use(middlewares.AuthMiddleware(true))
	users.UsersAuth(router.Group("/users"))

	exitVal := m.Run()

	models.CleanUpAfterTest()

	common.LogI.Println("Test main users services end with exit code", exitVal)
}

func createTest(asserts *assert.Assertions, testData *services.MockTests) *httptest.ResponseRecorder {
	body := testData.Body
	req, err := http.NewRequest(testData.Method, testData.Url, bytes.NewBufferString(body))

	asserts.NoError(err)

	testData.Init(req)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(testData.ResponseCode, w.Code, "Response Status - "+testData.Msg)

	return w
}

func TestUserRegister(t *testing.T) {
	asserts := assert.New(t)
	for _, testData := range services.MockTestsRegister {
		w := createTest(asserts, &testData)
		var jsonResp services.UserResponseMock
		err := json.Unmarshal(w.Body.Bytes(), &jsonResp)

		// common.LogI.Println("body", w.Body.String())

		if err != nil {
			panic("invalid json data")
		}
		tok := jsonResp.User.Token

		asserts.Regexp("(^[\\w-]*\\.[\\w-]*\\.[\\w-]*$)", tok, "Response Content - "+testData.Msg)
		// asserts.Regexp("(^[\\w-]*\\.[\\w-]*\\.[\\w-]*$)", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAzLTE1VDExOjA0OjU2LjkyMjE2MTIrMDc6MDAiLCJpZCI6MjgxfQ.Dbzjz5loj3X_lOG55gJtOXw2ENj2Re6sodnMmKPc-uc", "Response Content - "+testData.msg)
	}
}

func TestGetUsers(t *testing.T) {
	asserts := assert.New(t)
	for _, testData := range services.MockTestsGetUsers {
		w := createTest(asserts, &testData)
		var jsonResp services.UsersResponseMock
		err := json.Unmarshal(w.Body.Bytes(), &jsonResp)
		if err != nil {
			common.LogE.Println("json unmarshall error", err)
			panic("invalid json data")
		}

		// common.LogI.Println("jsonResp", jsonResp)
	}
}

func TestUserLogin(t *testing.T) {
	asserts := assert.New(t)
	for _, testData := range services.MockTestsLogin {
		w := createTest(asserts, &testData)
		var jsonResp services.UserResponseMock
		err := json.Unmarshal(w.Body.Bytes(), &jsonResp)

		// common.LogI.Println("jsonResp", jsonResp)

		asserts.Equal(testData.ResponseCode, w.Code, "Response status - "+testData.Msg)

		if err != nil {
			panic("invalid json data")
		}
		tok := jsonResp.User.Token

		asserts.Regexp("(^[\\w-]*\\.[\\w-]*\\.[\\w-]*$)", tok, "Response Content - "+testData.Msg)
	}
}
