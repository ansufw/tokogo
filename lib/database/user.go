package database

import (
	"fmt"

	"github.com/aysf/tokogo/config"
	jwtauth "github.com/aysf/tokogo/middlewares/jwtAuth"

	"github.com/aysf/tokogo/models"
)

func Signup(user *models.User) (interface{}, error) {
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

func LoginUser(user models.User) (interface{}, error) {
	var err error
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		fmt.Println("err 1")
		return nil, err
	}

	token, err := jwtauth.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println("err 2")
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		fmt.Println("err 3")
		return nil, err
	}
	return token, nil
}
