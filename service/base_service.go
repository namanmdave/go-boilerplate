package service

import (
	"go-boilerplate/store"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// BaseService contains common dependencies shared across all services
type BaseService struct {
	Store     *store.Store
	Logger    *zap.SugaredLogger
	Validator *validator.Validate
}

// NewBaseService creates a new base service with shared dependencies
func NewBaseService(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *BaseService {
	return &BaseService{
		Store:     store,
		Logger:    logger,
		Validator: validator,
	}
}
