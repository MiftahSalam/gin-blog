package serializers

import (
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (u *UserSerializer) Response() UserResponse {
	myUserModel := u.C.MustGet("user").(models.UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GetToken(myUserModel.ID),
	}

	return user
}

func (u *UserSerializer) Responses(users []models.UserModel) []UserResponse {
	var usersResponse []UserResponse
	for _, currentUser := range users {
		user := UserResponse{
			Username: currentUser.Username,
			Email:    currentUser.Email,
			Bio:      currentUser.Bio,
			Image:    currentUser.Image,
		}
		usersResponse = append(usersResponse, user)
	}
	// common.LogI.Println("usersResponse", usersResponse)

	return usersResponse
}
