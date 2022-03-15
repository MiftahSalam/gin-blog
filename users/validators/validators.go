package validators

import (
	"os"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
		Bio      string `form:"bio" json:"bio" binding:"max=1024"`
		Image    string `form:"image" json:"image" binding:"omitempty,url"`
	} `json:"user"`
	UserModel models.UserModel `json:"-"`
}

func (validator *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}

	validator.UserModel.Username = validator.User.Username
	validator.UserModel.Email = validator.User.Email
	validator.UserModel.Bio = validator.User.Bio
	validator.UserModel.SetPassword(validator.User.Password)
	if validator.User.Image != "" {
		validator.UserModel.Image = &validator.User.Image
	}

	return nil
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}

	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel models.UserModel) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.User.Username = userModel.Username
	userModelValidator.User.Email = userModel.Email
	userModelValidator.User.Bio = userModel.Bio
	userModelValidator.User.Password = os.Getenv("JWT_SECRET")

	if userModel.Image != nil {
		userModelValidator.User.Image = *userModel.Image
	}

	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (validator *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}

	validator.userModel.Email = validator.User.Email

	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}

	return loginValidator
}
