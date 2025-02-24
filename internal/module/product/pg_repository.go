package product

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type ProductRespository interface {
	GetList(ctx context.Context, categoryID *int) (models.GetListProductResp, error)
	GetProductByProductID(ctx context.Context, id int) (models.Product, error)
}
