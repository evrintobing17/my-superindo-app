package cart

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type CartUsecase interface {
	AddToCart(ctx context.Context, userID int, request models.AddToCardRequest) error
}
