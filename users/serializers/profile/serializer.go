package serializers

import (
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
)

type ProfileSerializer struct {
	C *gin.Context
	models.UserModel
}

type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"follwing"`
}

func (pSerializer *ProfileSerializer) Response() ProfileResponse {
	model := pSerializer.C.MustGet("my_user_model").(models.UserModel)
	profile := ProfileResponse{
		ID:        model.ID,
		Username:  model.Username,
		Bio:       model.Bio,
		Image:     model.Image,
		Following: model.IsFollowing(pSerializer.UserModel),
	}

	return profile
}
