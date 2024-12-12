package migration

import (
	"coding-chelleng/core/entities/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Car{}, &models.Order{})
}
