package usecase

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/cart"
)

type cartUsecase struct {
	repo cart.CartRepository
}

func NewCartUsecase(repo cart.CartRepository) cart.CartUsecase {
	return cartUsecase{repo: repo}
}

// AddToCart implements cart.CartUsecase.
func (c cartUsecase) AddToCart(ctx context.Context, userID int, request models.AddToCardRequest) error {
	ok := c.repo.IsCartExists(ctx, userID)
	if ok {
		err := c.repo.Insert(ctx, userID, request)
		return err
	}
	err := c.repo.Upsert(ctx, userID, request)

	return err
}
