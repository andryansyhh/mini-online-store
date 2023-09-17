package repository

import (
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"

	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

type TransactionRepository interface {
	CreateTransaction(req *models.Transaction) (*dto.TrxResponse, error)
}

func NewTransactionRepository(db *gorm.DB) Transaction {
	return Transaction{
		db: db,
	}
}

func (m *Transaction) CreateTransaction(req *models.Transaction) (*dto.TrxResponse, error) {
	if err := m.db.Table("transaction").Debug().Create(&req).Error; err != nil {
		return nil, err
	}

	var res dto.TrxResponse
	err := m.db.Table("transaction").Debug().
		Where("uuid = ?", req.Uuid).First(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
