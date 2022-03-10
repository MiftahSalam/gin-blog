package validators

import (
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		Bio      string `form:"bio" json:"bio" binding:"max=1024"`
		Image    string `form:"image" json:"image" binding:"omitempty,url"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (validator *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, validator)
	if err != nil {
		return err
	}

	validator.userModel.Username = validator.User.Username
	validator.userModel.Email = validator.User.Email
	validator.userModel.Bio = validator.User.Bio
	validator.userModel.SetPassword(validator.User.Password)
	if validator.User.Image != "" {
		validator.userModel.Image = &validator.User.Image
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
	userModelValidator.User.Password = common.RandString(10)

	if userModel.Image != nil {
		userModelValidator.User.Image = *userModel.Image
	}

	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
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
