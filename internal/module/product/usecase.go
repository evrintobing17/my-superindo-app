package product

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type ProductUsecase interface {
	GetListProduct(ctx context.Context, categoryID *int) (*models.GetListProductResp, error)
	GetProductByID(ctx context.Context, id int) (*models.Product, error)
}
