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
}

func AuthModdleware(autho401 bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UpdateContextUserModel(ctx, 0)
		tok, err := request.ParseFromRequest(ctx.Request, common.Auth2Extractor, func(t *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("JWT_SECRET"))
			return b, nil
		})
		if err != nil {
			if autho401 {
				ctx.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
			user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(ctx, user_id)
		}
	}
}
