package services

import (
	"fmt"
	"net/http"

	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
)

var MockTestsRegister = []struct {
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

var MockTestsGetUsers = []struct {
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
		"GET",
		"",
		http.StatusOK,
		"",
		"valid data end should return StatusOK",
	},
}

type UserResponseMock struct {
	User serializers.UserResponse
}

type UsersResponseMock struct {
	Users []serializers.UserResponse
}
