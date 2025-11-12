package handler

import (
	"go-boilerplate/service"
	"go-boilerplate/store"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// HandlerFactory is responsible for creating and managing all handler instances
// It uses a ServiceFactory to ensure services are properly shared
type HandlerFactory struct {
	baseHandler    *BaseHandler
	serviceFactory *service.ServiceFactory

	// Cache handler instances to reuse them
	userHandler    *UserHandler
	productHandler *ProductHandler
}

// NewHandlerFactory creates a new handler factory with shared dependencies
func NewHandlerFactory(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *HandlerFactory {
	baseHandler := NewBaseHandler(store, logger, validator)
	serviceFactory := service.NewServiceFactory(store, logger, validator)

	return &HandlerFactory{
		baseHandler:    baseHandler,
		serviceFactory: serviceFactory,
	}
}

// GetUserHandler returns a cached UserHandler instance or creates a new one
func (hf *HandlerFactory) GetUserHandler() *UserHandler {
	if hf.userHandler != nil {
		return hf.userHandler
	}

	hf.userHandler = NewUserHandler(
		hf.baseHandler,
		hf.serviceFactory.GetUserService(),
	)
	return hf.userHandler
}

// GetProductHandler returns a cached ProductHandler instance or creates a new one
func (hf *HandlerFactory) GetProductHandler() *ProductHandler {
	if hf.productHandler != nil {
		return hf.productHandler
	}

	hf.productHandler = NewProductHandler(
		hf.baseHandler,
		hf.serviceFactory.GetProductService(),
		hf.serviceFactory.GetUserService(), // Pass UserService as a dependency
	)
	return hf.productHandler
}
