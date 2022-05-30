package routers

import (
	"net/http/httptest"

	userServices "github.com/MiftahSalam/gin-blog/users/services"

	"github.com/stretchr/testify/assert"
)

type RouterMockTest struct {
	UserMockTest userServices.MockTests
	ResponseTest func(w *httptest.ResponseRecorder, a *assert.Assertions)
}
