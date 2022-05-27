package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	ArticleModels "github.com/MiftahSalam/gin-blog/articles/model"
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	common.LogI.Println("Test Main Article Services start")

	err := godotenv.Load("../../.env")
	if err != nil {
		common.LogE.Fatal("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}

	db := common.Init()

	ArticleModels.AutoMigrate()
	ArticleModels.Init(db)

	gin.SetMode(gin.TestMode)

	exitVal := m.Run()

	CleanUpAfterTest()
	sqlDB, err := db.DB()
	if err != nil {
		common.LogE.Fatal("get db instance error: ", err)
	}
	sqlDB.Close()

	os.Exit(exitVal)

	common.LogI.Println("Test Main Article Services end")
}

func TestCreateArticle(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleCreateTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)
			MockJSONPost(c, test.Data)

			ArticleCreate(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}

	// user, _ := c.Get("user")
	// common.LogI.Println("key user before", user)
	// for k := range c.Keys {
	// 	if k == "user" {
	// 		delete(c.Keys, k)
	// 		common.LogI.Println("ckeys", k)
	// 	}
	// }
	// user, _ = c.Get("user")
	// common.LogI.Println("key user after", user)
}

func TestGetArticlesFeed(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticlesFeedTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleFeed(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func MockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbyte, err := json.Marshal(content)
	if err != nil {
		common.LogE.Println("Cannot marshal json content")
		panic(err)
	}
	common.LogI.Println("content", string(jsonbyte))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbyte))
}

func InitTest() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	return c, w
}

func CleanUpAfterTest() {
	common.LogI.Println("clean up article services start")

	db := common.GetDB()
	for _, article := range ArticlesMock {
		createdArticleFromServices, err := ArticleModels.FindOneArticle(&article)
		if err != nil {
			common.LogE.Printf("cannot find article: %v with err %v", article.Title, err)
		}

		common.LogI.Println("clean up article tags", createdArticleFromServices.Tags)
		err = db.Unscoped().Model(&createdArticleFromServices).Association("Tags").Clear()
		if err != nil {
			common.LogE.Println("cannot delete article tags: ", err)
		}
	}

	ArticleModels.CleanUpAfterTest()

	common.LogI.Println("clean up article end")
}
