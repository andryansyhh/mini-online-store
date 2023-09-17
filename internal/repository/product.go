package repository

import (
	"mini-online-store/internal/domain/dto"

	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

type ProductRepository interface {
	GetListProducts() ([]dto.ProductResponse, error)
	GetListProductsByCategory(category string) ([]dto.ProductResponse, error)
	FindProductByUuid(uuid string) (*dto.ProductResponse, error)
}

func NewProductRepository(db *gorm.DB) Product {
	return Product{
		db: db,
	}
}

func (m *Product) GetListProducts() ([]dto.ProductResponse, error) {
	var res []dto.ProductResponse
	q := `product.uuid,product.name,product.qty,product.price, product.category_product_uuid,cp.name`
	err := m.db.Table("product").Select(q).Debug().
		Joins("JOIN category_product cp on cp.uuid = product.category_product_uuid").
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Product) GetListProductsByCategory(category string) ([]dto.ProductResponse, error) {
	var res []dto.ProductResponse
	q := `product.uuid,product.name,product.qty,product.price, product.category_product_uuid,cp.name`
	err := m.db.Table("product").Select(q).
		Joins("JOIN category_product cp on cp.uuid = product.category_product_uuid").
		Where("cp.name = ?", category).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Product) FindProductByUuid(uuid string) (*dto.ProductResponse, error) {
	var res dto.ProductResponse
	err := m.db.Table("product").Debug().
		Where("uuid = ?", uuid).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
