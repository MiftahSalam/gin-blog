package services

import (
	"errors"
	"net/http"

	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/middlewares"
	"github.com/MiftahSalam/gin-blog/users/models"
	profileSerializers "github.com/MiftahSalam/gin-blog/users/serializers/profile"
	userSerializers "github.com/MiftahSalam/gin-blog/users/serializers/user"
	"github.com/MiftahSalam/gin-blog/users/validators"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginValidation := validators.NewLoginValidator()
	if err := loginValidation.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	userModel, err := models.FindOneUser(&models.UserModel{Email: loginValidation.User.Email})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered or invalid email")))
		return
	}
	if userModel.CheckPassword(loginValidation.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("invalid password")))
		return
	}

	middlewares.UpdateContextUserModel(c, userModel.ID)
	serializer := userSerializers.UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func Register(c *gin.Context) {
	userValidation := validators.NewUserModelValidator()
	if err := userValidation.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModel, _ := models.FindOneUser("username = ?", userValidation.User.Username)
	if userModel.Username == "" {
		if err := models.SaveOne(&userValidation.UserModel); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
			return
		}

		c.Set("user", userValidation.UserModel)
		serializer := userSerializers.UserSerializer{C: c}
		c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
	} else {
		c.JSON(http.StatusBadRequest, common.NewError("Validation", errors.New("user already exist")))
		return
	}
	// common.LogI.Println("Find existing user", userModel.Username)
}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		common.LogI.Println("get users error", err)
		c.JSON(http.StatusInternalServerError, common.NewError("database error", err))
	}
	// common.LogI.Println("users", users)

	serializer := userSerializers.UserSerializer{C: c}
	c.JSON(http.StatusOK, gin.H{"users": serializer.Responses(users)})
}

func UpdateUser(c *gin.Context) {
	userValidation := validators.NewUserModelValidator()
	if err := userValidation.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	if user_id := c.MustGet("user_id").(uint); user_id > 0 {
		// common.LogI.Printf("user_id ctx %v, type %T", user_id, user_id)
		userModel, _ := models.FindOneUser("id = ?", user_id)

		userUpdate := userValidation.UserModel

		if userModel.Username != "" {
			if err := userModel.Update(&userUpdate); err == nil {
				// common.LogI.Printf("userValidation %v", userValidation.User)
				c.Set("user", userModel)
				serializer := userSerializers.UserSerializer{C: c}
				c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
			} else {
				common.LogE.Printf("error update")
				c.JSON(http.StatusInternalServerError, common.NewError("Server error", err))
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, common.NewError("Request error", errors.New("cannot find user")))
			return
		}
	} else {
		common.LogI.Println("user_id not exist")
		c.JSON(http.StatusUnauthorized, common.NewError("Request error", errors.New("not logged in")))
		return
	}
}

//todo delete user service

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	userModel, err := models.FindOneUser("username = ? ", username)

	common.LogI.Println("username", username)
	common.LogI.Println("userModel", userModel)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("invalid username")))
		return
	}

	profileSerializer := profileSerializers.ProfileSerializer{C: c, UserModel: userModel}
	c.JSON(http.StatusOK, gin.H{"profile": profileSerializer.Response()})
}
