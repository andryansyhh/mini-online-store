package migrations

import (
	"mini-online-store/internal/domain/models"

	"gorm.io/gorm"
)

func CreateTrx(db *gorm.DB) error {
	return db.AutoMigrate(&models.Transaction{})
}
