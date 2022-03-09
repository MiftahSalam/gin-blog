package services

import (
	"github.com/MiftahSalam/gin-blog/common"
	"github.com/MiftahSalam/gin-blog/users/models"
)

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
