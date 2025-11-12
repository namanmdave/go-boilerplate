package handler

import (
	"go-boilerplate/service"
)

// ProductHandler handles HTTP requests related to products
// It demonstrates how to use multiple services (in this case, both ProductService and UserService)
type ProductHandler struct {
	*BaseHandler

	ProductService *service.ProductService
	UserService    *service.UserService
}

// NewProductHandler creates a new product handler with base handler and services
func NewProductHandler(
	baseHandler *BaseHandler,
	productService *service.ProductService,
	userService *service.UserService,
) *ProductHandler {
	return &ProductHandler{
		BaseHandler:    baseHandler,
		ProductService: productService,
		UserService:    userService,
	}
}

// Example handler methods - add your product endpoints here:
// func (ph *ProductHandler) GetProduct(c *gin.Context) {
//     productID := c.Param("id")
//     // Use ph.ProductService to get product
//     // Use ph.UserService if needed
//     // Use ph.Logger for logging
//     // Use ph.Validator for validation
// }
//
// func (ph *ProductHandler) CreateProduct(c *gin.Context) {
//     var req CreateProductRequest
//     if err := c.ShouldBindJSON(&req); err != nil {
//         ph.Logger.Errorw("failed to bind request", "error", err)
//         c.JSON(400, gin.H{"error": "invalid request"})
//         return
//     }
//
//     if err := ph.Validator.Struct(req); err != nil {
//         ph.Logger.Errorw("validation failed", "error", err)
//         c.JSON(400, gin.H{"error": "validation failed"})
//         return
//     }
//
//     // Use the product service which internally uses user service
//     // product, err := ph.ProductService.CreateProduct(c.Request.Context(), &req)
// }
