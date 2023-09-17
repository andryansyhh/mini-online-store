package usecase

import (
	"errors"
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"
	"mini-online-store/internal/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	CreateTransaction(req *models.Transaction) (*dto.TrxResponse, error)
}

type transactionUsecase struct {
	productRepository     repository.ProductRepository
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
	}
}

func (m *transactionUsecase) CreateTransaction(req *models.Transaction) (*dto.TrxResponse, error) {
	product, err := m.productRepository.FindProductByUuid(req.ProductUuid)
	if err != nil {
		return nil, errors.New("failed get product")
	}

	if product.Uuid == "" {
		return nil, errors.New("product not found")
	}
	if req.Price < product.Price {
		return nil, errors.New("not enough amount")

	}

	req.Uuid = uuid.New().String()
	req.CreatedAt = time.Now()
	res, err := m.transactionRepository.CreateTransaction(req)
	if err != nil {
		return nil, errors.New("failed create transaction")
	}

	return res, nil
}
