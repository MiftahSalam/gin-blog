package services

import (
	"errors"
	"net/http"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/MiftahSalam/gin-blog/users/models"
	serializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
	"github.com/MiftahSalam/gin-blog/users/validators"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginValidation := validators.NewLoginValidator()
	if err := loginValidation.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModel, err := models.FindOneUser(&models.UserModel{Email: loginValidation.User.Email})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered or invalid email")))
	}
	if userModel.CheckPassword(loginValidation.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("invalid password")))
	}

	middlewares.UpdateContextUserModel(c, userModel.ID)
	serializer := serializers.UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

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
	c.JSON(http.StatusOK, gin.H{"users": serializer.Responses(users)})
}
