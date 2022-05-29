package services

import (
	"encoding/json"
	"io/ioutil"
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

func TestArticleRetrieve(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleRetrieveTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleRetrieve(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestArticleUpdate(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleUpdate {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)
			MockJSONPost(c, test.Data)

			ArticleUpdate(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestArticleFavorite(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleFavoriteTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleFavorite(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestArticleUnFavorite(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleUnFavoriteTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleUnFavorite(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestCreateArticleComment(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleCommentCreateTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)
			MockJSONPost(c, test.Data)

			ArticleCommentCreate(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestListArticleComment(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleCommentListTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleCommentList(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestArticleCommenteDelete(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleCommentDeleteTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleCommentDelete(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestTagList(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleCommentListTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()

			TagList(c)

			asserts.Equal(http.StatusOK, w.Code)

			response_body, _ := ioutil.ReadAll(w.Body)

			common.LogI.Println("response_body", string(response_body))

			var jsonResp TagsResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				common.LogE.Println("Cannot umarshal json content with error: ", err)
			}
			asserts.NoError(err)

			common.LogI.Println("jsonResp", jsonResp)

			allTags := append(ArticleModels.TagsMock, tagsMockUpdate...)
			asserts.Equal(uint(len(allTags)), uint(len(jsonResp.Tags)))
			asserts.Equal(allTags, jsonResp.Tags)
		})
	}
}

func TestArticleDelete(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockArticleDeleteTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			ArticleDelete(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
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
	for _, tag := range tagsMockUpdate {
		common.LogI.Println("clean up tag", tag)

		var tagModel ArticleModels.TagModel
		err := db.Unscoped().Delete(&tagModel, ArticleModels.TagModel{
			Tag: tag,
		}).Error
		if err != nil {
			common.LogE.Println("cannot delete tag: ", err)
		}
	}

	ArticleModels.CleanUpAfterTest()

	common.LogI.Println("clean up article end")
}
