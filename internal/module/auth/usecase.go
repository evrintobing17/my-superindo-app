package auth

import (
	"context"

	"github.com/evrintobing17/my-superindo-app/internal/models"
)

type AuthUsecase interface {
	SignUp(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (string, error)
}
