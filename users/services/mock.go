package services

import (
	"fmt"
	"net/http"

	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
)

type MockTests struct {
	TestName        string
	Init            func(*http.Request)
	Url             string
	Method          string
	Body            string
	ResponseCode    int
	ResponsePattern string
	Msg             string
}

var MockTestsLogin = []MockTests{
	{
		"no error: Login Test",
		func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		"/users/login",
		"POST",
		fmt.Sprintf(`{"user":{"email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1),
		http.StatusOK,
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","bio":"", "image":"null", "token":"([a-zA-Z0-9-_.])"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		"valid data end should return StatusOk",
	},
}

var MockTestsRegister = []MockTests{
	{
		"no error: Register Test",
		func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		"/users/",
		"POST",
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		http.StatusCreated,
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","bio":"", "image":"null", "token":"([a-zA-Z0-9-_.])"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		"valid data end should return StatusCreated",
	},
	{
		"error unproccesed data (no header content-type tag): Register Test",
		func(req *http.Request) {},
		"/users/",
		"POST",
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		http.StatusUnprocessableEntity,
		"",
		"valid data end should return StatusUnprocessableEntity",
	},
	{
		"error user already exist: Register Test",
		func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		"/users/",
		"POST",
		fmt.Sprintf(`{"user":{"username":"user%v","email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1, models.UserMockNumber+1),
		http.StatusBadRequest,
		"",
		"valid data end should return Bad request",
	},
}

var MockTestsGetUsers = []MockTests{
	{
		"no error: Get Users Test",
		func(req *http.Request) {
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTA0LTAzVDExOjE1OjE0Ljk0Nzg2NjIrMDc6MDAiLCJpZCI6NDg1fQ.IBzRW627TBLpYFFj2-6DDaXcPBkv4XW5dtMuSr6aohY")
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
