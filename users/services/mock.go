package services

import (
	"fmt"
	"net/http"

	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
)

var MockTestsLogin = []struct {
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
		"/users/login",
		"POST",
		fmt.Sprintf(`{"user":{"email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1),
		http.StatusOK,
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","bio":"", "image":"null", "token":"([a-zA-Z0-9-_.])"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		"valid data end should return StatusOk",
	},
}

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
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
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
		func(req *http.Request) {
			req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTA0LTAzVDExOjE1OjE0Ljk0Nzg2NjIrMDc6MDAiLCJpZCI6NDg1fQ.IBzRW627TBLpYFFj2-6DDaXcPBkv4XW5dtMuSr6aohY")
		},
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
