package services

import (
	"net/http"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
	"github.com/MiftahSalam/gin-blog/users/validators"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userValidation := validators.NewUserModelValidator()
	if err := userValidation.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := models.SaveOne(&userValidation.UserModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	c.Set("user", userValidation.UserModel)
	serializer := serializers.UserSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		common.LogI.Println("get users error", err)
		c.JSON(http.StatusInternalServerError, common.NewError("database error", err))
	}
	// common.LogI.Println("users", users)

	serializer := serializers.UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Responses(users)})
}
