package service

import (
	"context"
	"fmt"
	"go-boilerplate/dto"
	"go-boilerplate/store"
	db "go-boilerplate/store/sqlc"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// UserService handles business logic for user operations
type UserService struct {
	*BaseService
}

// NewUserService creates a new user service with base dependencies
func NewUserService(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *UserService {
	return &UserService{
		BaseService: NewBaseService(store, logger, validator),
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *dto.CreateUserRequest) (db.CreateUserRow, error) {
	dbUser, err := us.Store.Querier.CreateUser(ctx, db.CreateUserParams{
		Email: user.Email,
		Name:  user.Username,
	})
	if err != nil {
		return db.CreateUserRow{}, fmt.Errorf("failed to create user: %w", err)
	}

	return dbUser, nil
}
