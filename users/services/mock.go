package services

import (
	"fmt"
	"net/http"

	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
)

var MockTests = []struct {
	Init            func(*http.Request)
	Url             string
	Method          string
	Body            string
	ResponseCode    int
	ResponsePattern string
	Msg             string
}{
	{
		func(req *http.Request) {},
		"/users/",
		"POST",
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","password":"123456"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		http.StatusCreated,
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","bio":"", "image":"null", "token":"([a-zA-Z0-9-_.])"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		"valid data end should return StatusCreated",
	},
}

type UserResponseMock struct {
	User serializers.UserResponse
}
