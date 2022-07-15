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
	logged_user_id_str, _ := pSerializer.C.Get("user_id")
	logged_user_id := logged_user_id_str.(uint)
	// common.LogI.Println("logged_user_id_str", logged_user_id_str)
	// common.LogI.Println("logged_user_id", logged_user_id)
	var isFollowing bool = false
	if logged_user_id > 0 {
		currentUser := pSerializer.C.MustGet("user").(models.UserModel)
		isFollowing = currentUser.IsFollowing(pSerializer.UserModel)
		// common.LogI.Printf("%v isFollowing %v %s", logged_user_id, isFollowing, pSerializer.Username)
	}
	profile := ProfileResponse{
		ID:        pSerializer.ID,
		Username:  pSerializer.Username,
		Bio:       pSerializer.Bio,
		Image:     pSerializer.Image,
		Following: isFollowing,
	}

	return profile
}
