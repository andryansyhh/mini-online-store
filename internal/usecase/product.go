package usecase

import (
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/repository"
)

type ProductUsecase interface {
	GetAllProduct() ([]dto.ProductResponse, error)
	GetProductByCategory(category string) ([]dto.ProductResponse, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepo,
	}
}

func (p *productUsecase) GetAllProduct() ([]dto.ProductResponse, error) {
	res, err := p.productRepository.GetListProducts()
	if err != nil {
		return nil, err
	}

	var resp []dto.ProductResponse
	for _, v := range res {
		resProduct := dto.ProductResponse{
			Uuid:                v.Uuid,
			Name:                v.Name,
			CategoryProductName: v.CategoryProductName,
			Qty:                 v.Qty,
			Price:               v.Price,
		}

		resp = append(resp, resProduct)
	}

	return resp, nil
}

func (p *productUsecase) GetProductByCategory(category string) ([]dto.ProductResponse, error) {
	res, err := p.productRepository.GetListProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	var resp []dto.ProductResponse
	for _, v := range res {
		resProduct := dto.ProductResponse{
			Uuid:                v.Uuid,
			Name:                v.Name,
			CategoryProductName: v.CategoryProductName,
			Qty:                 v.Qty,
			Price:               v.Price,
		}

		resp = append(resp, resProduct)
	}

	return resp, nil
}
