package service

import (
	"go-boilerplate/store"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ProductService handles business logic for product operations
// It demonstrates how to use multiple services (in this case, UserService)
type ProductService struct {
	*BaseService

	// Dependencies on other services
	UserService *UserService
}

// NewProductService creates a new product service with base dependencies and other services
func NewProductService(
	store *store.Store,
	logger *zap.SugaredLogger,
	validator *validator.Validate,
	userService *UserService,
) *ProductService {
	return &ProductService{
		BaseService: NewBaseService(store, logger, validator),
		UserService: userService,
	}
}

// Example methods - add your product business logic here:
// func (ps *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
//     if err := ps.Validator.Struct(product); err != nil {
//         ps.Logger.Errorw("validation failed", "error", err)
//         return err
//     }
//
//     // You can use UserService here if needed
//     user, err := ps.UserService.GetUserByID(ctx, product.OwnerID)
//     if err != nil {
//         ps.Logger.Errorw("failed to get product owner", "error", err)
//         return err
//     }
//
//     // Now create the product
//     // return ps.Store.CreateProduct(ctx, product)
// }
