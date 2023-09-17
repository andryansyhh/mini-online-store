package usecase

import (
	"errors"
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/repository"

	"github.com/google/uuid"
)

type ShoppingCartUsecase interface {
	AddToCart(req *dto.AddToCart) error
	ListItemInCart(user_uuid string) ([]dto.CartResponse, error)
	DeleteCartItem(cart_uuid, user_uuid string) error
}

type shoppingCartUsecase struct {
	// userRepository         repository.UserRepository
	productRepository      repository.ProductRepository
	shoppingCartRepository repository.ShoppingCartRepository
}

func NewShoppingCartUsecase(shoppingCartRepository repository.ShoppingCartRepository, productRepository repository.ProductRepository) ShoppingCartUsecase {
	return &shoppingCartUsecase{
		shoppingCartRepository: shoppingCartRepository,
		productRepository:      productRepository,
	}
}

func (u *shoppingCartUsecase) AddToCart(req *dto.AddToCart) error {
	product, err := u.productRepository.FindProductByUuid(req.ProductUuid)
	if err != nil {
		return errors.New("failed get product")
	}

	if product.Uuid == "" {
		return errors.New("product not found")
	}

	req.Uuid = uuid.New().String()
	if err := u.shoppingCartRepository.AddToCart(req); err != nil {
		return errors.New("failed add to cart")
	}

	return nil
}

func (u *shoppingCartUsecase) ListItemInCart(user_uuid string) ([]dto.CartResponse, error) {
	res, err := u.shoppingCartRepository.GetListCart(user_uuid)
	if err != nil {
		return nil, errors.New("failed get product")
	}

	var resp []dto.CartResponse
	for _, v := range res {
		resProduct := dto.CartResponse{
			Uuid:        v.Uuid,
			UserUuid:    v.UserUuid,
			ProductName: v.ProductName,
			Qty:         v.Qty,
			Price:       v.Price,
		}

		resp = append(resp, resProduct)
	}

	return resp, nil
}

func (u *shoppingCartUsecase) DeleteCartItem(cart_uuid, user_uuid string) error {
	res, err := u.shoppingCartRepository.GetListCartByUuid(user_uuid, cart_uuid)
	if err != nil {
		return errors.New("failed get product")
	}

	if res.Uuid == "" {
		return errors.New("cart not found")
	}

	err = u.shoppingCartRepository.DeleteProductInCart(cart_uuid)
	if err != nil {
		return errors.New("failed delete product")
	}

	return nil
}
