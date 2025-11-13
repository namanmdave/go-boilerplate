package service

import (
	"go-boilerplate/store"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ServiceFactory is responsible for creating and managing all service instances
type ServiceFactory struct {
	store     *store.Store
	logger    *zap.SugaredLogger
	validator *validator.Validate

	// Cache service instances to reuse them
	userService    *UserService
	productService *ProductService
	messageService *MessageService
	groupService   *GroupService
}

// NewServiceFactory creates a new service factory with shared dependencies
func NewServiceFactory(store *store.Store, logger *zap.SugaredLogger, validator *validator.Validate) *ServiceFactory {
	return &ServiceFactory{
		store:     store,
		logger:    logger,
		validator: validator,
	}
}

// GetUserService returns a cached UserService instance or creates a new one
func (sf *ServiceFactory) GetUserService() *UserService {
	if sf.userService != nil {
		return sf.userService
	}

	sf.userService = NewUserService(sf.store, sf.logger, sf.validator)
	return sf.userService
}

// GetProductService returns a cached ProductService instance or creates a new one
func (sf *ServiceFactory) GetProductService() *ProductService {
	if sf.productService != nil {
		return sf.productService
	}

	// ProductService depends on UserService
	// The factory handles this dependency automatically
	sf.productService = NewProductService(
		sf.store,
		sf.logger,
		sf.validator,
		sf.GetUserService(), // Pass other services here
	)
	return sf.productService
}

func (sf *ServiceFactory) GetMessageService() *MessageService {
	if sf.messageService != nil {
		return sf.messageService
	}

	sf.messageService = NewMessageService(sf.store, sf.logger, sf.validator)
	return sf.messageService
}

func (sf *ServiceFactory) GetGroupService() *GroupService {
	if sf.groupService != nil {
		return sf.groupService
	}

	sf.groupService = NewGroupService(sf.store, sf.logger, sf.validator)
	return sf.groupService
}
