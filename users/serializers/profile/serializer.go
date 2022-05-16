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
	Following bool    `json:"following"`
}

func (pSerializer *ProfileSerializer) Response() ProfileResponse {
	model := pSerializer.C.MustGet("user").(models.UserModel)
	profile := ProfileResponse{
		ID:        pSerializer.ID,
		Username:  pSerializer.Username,
		Bio:       pSerializer.Bio,
		Image:     pSerializer.Image,
		Following: model.IsFollowing(pSerializer.UserModel),
	}

	return profile
}
