package handler

import (
	"go-boilerplate/store"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// BaseHandler contains common dependencies shared across all handlers
type BaseHandler struct {
	Store     *store.Store
	Logger    *zap.SugaredLogger
	Validator *validator.Validate
}

// NewBaseHandler creates a new base handler with shared dependencies
func NewBaseHandler(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *BaseHandler {
	return &BaseHandler{
		Store:     store,
		Logger:    logger,
		Validator: validator,
	}
}
