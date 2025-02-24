package auth

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type AuthRepository interface {
	Login(ctx context.Context, email string) (*models.User, error)
	SignUp(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, userID int) error
}
