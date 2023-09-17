package repository

import (
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	db *gorm.DB
}

type ShoppingCartRepository interface {
	AddToCart(addToCart *dto.AddToCart) error
	GetListCart(user_uuid string) ([]dto.CartResponse, error)
	DeleteProductInCart(cart_uuid string) error
	GetListCartByUuid(cart_uuid, user_uuid string) (*dto.CartResponse, error)
}

func NewShoppingCartRepository(db *gorm.DB) ShoppingCart {
	return ShoppingCart{
		db: db,
	}
}

func (m *ShoppingCart) AddToCart(addToCart *dto.AddToCart) error {
	if err := m.db.Table("shopping_cart").Debug().Create(&addToCart).Error; err != nil {
		return err
	}
	return nil
}

func (m *ShoppingCart) GetListCart(user_uuid string) ([]dto.CartResponse, error) {
	var res []dto.CartResponse
	q := `shopping_cart.uuid,shopping_cart.user_uuid,p.name,p.price,shopping_cart.qty `
	err := m.db.Table("shopping_cart").Select(q).Debug().
		Joins("JOIN product p on p.uuid = shopping_cart.product_uuid").
		Where("shopping_cart.user_uuid = ?", user_uuid).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *ShoppingCart) GetListCartByUuid(user_uuid, cart_uuid string) (*dto.CartResponse, error) {
	var res dto.CartResponse
	q := `shopping_cart.uuid,shopping_cart.user_uuid,p.name,p.price,shopping_cart.qty `
	err := m.db.Table("shopping_cart").Select(q).Debug().
		Joins("JOIN product p on p.uuid = shopping_cart.product_uuid").
		Where("shopping_cart.user_uuid = ? and shopping_cart.uuid = ?", user_uuid, cart_uuid).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *ShoppingCart) DeleteProductInCart(cart_uuid string) error {
	var res models.ShoppingCart
	err := m.db.Table("shopping_cart").
		Where("uuid = ?", cart_uuid).Delete(&res).Error
	if err != nil {
		return err
	}
	return nil
}
