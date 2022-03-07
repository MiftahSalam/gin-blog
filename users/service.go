package users

import (
	"errors"

	"github.com/MiftahSalam/gin-blog/common"
	"golang.org/x/crypto/bcrypt"
)

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
func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel

	err := db.Where(condition).First(&model).Error

	return model, err
}
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error

	return err
}

func (u UserModel) following(fu UserModel) error {
	db := common.GetDB()
	var follow FollowModel
	err := db.FirstOrCreate(&follow, &FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).Error

	return err
}

func (u UserModel) isFollowing(fu UserModel) bool {
	db := common.GetDB()
	var follow FollowModel

	db.Where(FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).First(&follow)

	return follow.ID != 0

}

func (u UserModel) unFollow(fu UserModel) error {
	db := common.GetDB()
	err := db.Where(FollowModel{
		FollowingID:  fu.ID,
		FollowedByID: u.ID,
	}).Delete(FollowModel{}).Error

	return err
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should no be empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)

	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)

	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
