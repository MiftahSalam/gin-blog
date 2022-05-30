package routers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	articleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	common.LogI.Println("Test main article router start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}

	db := common.Init()
	articleModels.AutoMigrate()
	articleModels.Init(db)

	gin.SetMode(gin.TestMode)

	router = gin.New()
	router.Use(middlewares.AuthMiddleware(true))
	Articles(router.Group("/article"))

	exitVal := m.Run()

	articleModels.CleanUpAfterTest()
	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	os.Exit(exitVal)

	common.LogI.Println("Test main article router end")
}

func createTest(asserts *assert.Assertions, testData *RouterMockTest) *httptest.ResponseRecorder {
	body := testData.UserMockTest.Body
	req, err := http.NewRequest(testData.UserMockTest.Method, testData.UserMockTest.Url, bytes.NewBufferString(body))

	asserts.NoError(err)

	testData.UserMockTest.Init(req)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(testData.UserMockTest.ResponseCode, w.Code, "Response Status - "+testData.UserMockTest.Msg)

	return w
}

func TestGetArticle(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockGetArticle {
		t.Run(test.UserMockTest.TestName, func(t *testing.T) {
			w := createTest(asserts, &test)

			asserts.Equal(test.UserMockTest.ResponseCode, w.Code)

			test.ResponseTest(w, asserts)
		})
	}

}
