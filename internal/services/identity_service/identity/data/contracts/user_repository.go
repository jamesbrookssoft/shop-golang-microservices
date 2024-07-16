package contracts

import (
	"context"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/models"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
}
