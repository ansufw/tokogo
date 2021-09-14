package database

import (
	"github.com/aysf/tokogo/config"
	"github.com/aysf/tokogo/models"
)

func GenerateCart(cartID models.Cart) (interface{}, error) {

	if err := config.DB.Create(&cartID).Error; err != nil {
		return nil, err
	}
	return cartID, nil
}
