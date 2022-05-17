package services

import (
	"fmt"
	"net/http"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	userProfileSerializers "github.com/MiftahSalam/gin-blog/users/serializers/profile"
	userSerializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
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
		"error bad request (no header content-type tag): Login Test",
		func(req *http.Request) {},
		"/users/login",
		"POST",
		fmt.Sprintf(`{"user":{"email":"%v@gmail.com","password":"12345678"}}`, models.UserMockNumber+1),
		http.StatusBadRequest,
		"{}",
		"valid data end should return StatusBadRequest",
	},
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

var MockTestsUpdateUser = []MockTests{
	{
		"no error: Update Test",
		func(req *http.Request) {
			// common.LogI.Println("username", models.UsersMock[models.UserMockNumber-int(models.CurrentRecordCount)-1].Username)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(models.UsersMock[models.UserMockNumber-int(models.CurrentRecordCount)-1].ID)))
		},
		"/users/",
		"PUT",
		`{"user":{"username":"userUpdated","email":"Updated@gmail.com","bio":"bioUpdated","password":"passUpdated"}}`,
		http.StatusOK,
		`{"user":{"username":"userUpdated","email":"Updated@gmail.com","bio":"bioUpdated"}}`,
		"valid data end should return StatusOk",
	},
	{
		"error unproccesed data (no header content-type tag): Update Test",
		func(req *http.Request) {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(models.UsersMock[models.UserMockNumber-int(models.CurrentRecordCount)-1].ID)))
		},
		"/users/",
		"PUT",
		`{"user":{"username":"userUpdated","email":"Updated@gmail.com","bio":"bioUpdated","password":"passUpdated"}}`,
		http.StatusBadRequest,
		"{}",
		"valid data end should return StatusUnprocessableEntity",
	},
	{
		"error unauthorized (no header authorization): Update Test",
		func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		"/users/",
		"PUT",
		`{"user":{"username":"userUpdated","email":"Updated@gmail.com","bio":"bioUpdated","password":"passUpdated"}}`,
		http.StatusUnauthorized,
		"{}",
		"valid data end should return StatusUnauthorized",
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
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(models.UsersMock[models.UserMockNumber-int(models.CurrentRecordCount)-1].ID)))
		},
		"/users/",
		"GET",
		"",
		http.StatusOK,
		"",
		"valid data end should return StatusOK",
	},
	{
		"error unauthorized: token not exist",
		func(req *http.Request) {
			// req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTA0LTAzVDExOjE1OjE0Ljk0Nzg2NjIrMDc6MDAiLCJpZCI6NDg1fQ.IBzRW627TBLpYFFj2-6DDaXcPBkv4XW5dtMuSr6aohY")
		},
		"/users/",
		"GET",
		"",
		http.StatusUnauthorized,
		"",
		"invalid data end should return StatusUnauthorized",
	},
	{
		"error unauthorized: token expired",
		func(req *http.Request) {
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTI0Mzg3MzksImlkIjo2MTR9.kaUsapW0gBxmBcu5C3LsXbbeMnMfAOhMD4Ri7jc9ZlI")
		},
		"/users/",
		"GET",
		"",
		http.StatusUnauthorized,
		"",
		"invalid data end should return StatusUnauthorized",
	},
}

var MockTestsGetUserProfile = []MockTests{
	{
		"no error: Get User Profile",
		func(req *http.Request) {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", common.GetToken(models.UsersMock[models.UserMockNumber-int(models.CurrentRecordCount)-1].ID)))
		},
		fmt.Sprintf("/profile/user%v", 0),
		"GET",
		"",
		http.StatusOK,
		fmt.Sprintf(`{"profile":{"username":"user%v","bio":"","image":"null","following":"false"}}`, 0),
		"valid data end should return StatusOK",
	},
	{
		"error unauthorized: token not exist. Get User Profile",
		func(req *http.Request) {},
		fmt.Sprintf("/profile/user%v", 0),
		"GET",
		"",
		http.StatusUnauthorized,
		"{}",
		"valid data end should return StatusOK",
	},
}

type UserResponseMock struct {
	User userSerializers.UserResponse
}

type UsersResponseMock struct {
	Users []userSerializers.UserResponse
}

type UserProfileResponseMock struct {
	Profile userProfileSerializers.ProfileResponse
}
