package usecase

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/product"
)

type productUsecase struct {
	repository product.ProductRespository
}

func NewProductUsecase(repository product.ProductRespository) product.ProductUsecase {
	return &productUsecase{
		repository: repository,
	}
}

// GetListProduct implements product.ProductUsecase.
func (p *productUsecase) GetListProduct(ctx context.Context, categoryID *int) (*models.GetListProductResp, error) {
	product, err := p.repository.GetList(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// GetProductByID implements product.ProductUsecase.
func (p *productUsecase) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	product, err := p.repository.GetProductByProductID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
