package database

import (
	"github.com/aysf/tokogo/config"
	"github.com/aysf/tokogo/models"
)

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(UserID int) (interface{}, error) {
	var user models.User

	if err := config.DB.Find(&user, UserID).Error; err != nil {
		return nil, err
	}
	return user, nil
}
