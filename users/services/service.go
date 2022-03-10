package services

import (
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
)

func GetUsers() ([]models.UserModel, error) {
	db := common.GetDB()
	var users []models.UserModel

	err := db.Find(&users).Error

	return users, err
}

func FindOneUser(condition interface{}) (models.UserModel, error) {
	db := common.GetDB()
	var model models.UserModel

	err := db.Where(condition).First(&model).Error

	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error

	return err
}

func DeleteOneUsers(data interface{}) error {
	db := common.GetDB()
	err := db.Unscoped().Delete(models.UserModel{}, data).Error

	return err
}
