package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	//ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    int
	ProductID int
}
