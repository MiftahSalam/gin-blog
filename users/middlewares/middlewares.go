package middlewares

import (
	"net/http"
	"os"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func UpdateContextUserModel(c *gin.Context, user_id uint) {
	var userModel models.UserModel

	if user_id != 0 {
		db := common.GetDB()
		db.First(&userModel, user_id)
	}
	c.Set("user_id", user_id)
	c.Set("user", userModel)

	common.LogI.Println("current_user_id", user_id)
}

func AuthMiddleware(autho401 bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		common.LogI.Println("autho401", autho401)
		// if autho401 {
		UpdateContextUserModel(ctx, 0)
		tok, err := request.ParseFromRequest(ctx.Request, common.Auth2Extractor, func(t *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("JWT_SECRET"))
			return b, nil
		})

		if err != nil {
			common.LogE.Println("err", err)
			if autho401 {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				common.LogI.Println("no tok check")
				return
			}
			return
		}

		if claims, ok := tok.Claims.(jwt.MapClaims); ok {
			// common.LogI.Println("cek tok exp", claims["exp"])

			user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(ctx, user_id)
		}
		// }
	}
}
