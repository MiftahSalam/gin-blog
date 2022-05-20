package models

import (
	"errors"

	"github.com/MiftahSalam/gin-blog/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null;->:false;<-:create"`
}

type FollowModel struct {
	gorm.Model
	Following    UserModel
	FollowingID  uint
	FollowedBy   UserModel
	FollowedByID uint
}

func (u FollowModel) TableName() string {
	return "follows"
}

func (u UserModel) TableName() string {
	return "users"
}

func AuthoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FollowModel{})

}

func (u *UserModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(u).Updates(data).Error

	return err
}

func (u UserModel) GetFollowing() []UserModel {
	db := common.GetDB()
	tx := db.Begin()
	var follows []FollowModel
	var followings []UserModel

	tx.Where(FollowModel{
		FollowedByID: u.ID,
	}).Find(&follows)

	for _, follow := range follows {
		var userModel UserModel

		tx.Model(&follow).Association("Following").Find(&userModel)
		followings = append(followings, userModel)
	}
	tx.Commit()

	return followings
}

func (u UserModel) Following(fu UserModel) error {
	db := common.GetDB()
	var follow FollowModel
	err := db.FirstOrCreate(&follow, &FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).Error

	return err
}

func (u UserModel) IsFollowing(fu UserModel) bool {
	db := common.GetDB()
	var follow FollowModel

	db.Where(FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).First(&follow)

	return follow.ID != 0

}

func (u UserModel) UnFollow(fu UserModel) error {
	db := common.GetDB()
	err := db.Unscoped().Where(FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).Delete(FollowModel{}).Error

	return err
}

func (u *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should no be empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)

	return nil
}

func (u *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)

	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func GetUsers() ([]UserModel, error) {
	db := common.GetDB()
	var users []UserModel

	err := db.Find(&users).Error

	return users, err
}

func FindOneUser(condition interface{}, args ...interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel

	err := db.Where(condition, args...).First(&model).Error

	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error

	return err
}

func DeleteOneUsers(data interface{}) error {
	user, ok := data.(UserModel)
	if !ok {
		return errors.New("invalid user model input")
	}

	db := common.GetDB()

	db.Unscoped().Delete(FollowModel{}, &FollowModel{
		FollowedByID: user.ID,
	})
	db.Unscoped().Delete(FollowModel{}, &FollowModel{
		FollowingID: user.ID,
	})

	err := db.Unscoped().Delete(UserModel{}, data).Error

	return err
}
