package cart

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type CartRepository interface {
	IsCartExists(ctx context.Context, userID int) bool
	Insert(ctx context.Context, userID int, request models.AddToCardRequest) error
	Upsert(ctx context.Context, userID int, request models.AddToCardRequest) error
}
