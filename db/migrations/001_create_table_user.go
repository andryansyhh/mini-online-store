package migrations

import (
	"mini-online-store/internal/domain/models"

	"gorm.io/gorm"
)

func CreateUserTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
