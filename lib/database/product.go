package database

import (
	"github.com/aysf/tokogo/config"
	"github.com/aysf/tokogo/models"
)

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
