/*
USING MULTIPLE SERVICES IN A SINGLE SERVICE
==============================================

This document explains the pattern for using multiple services within a single service.

SCENARIO: ProductService needs to use UserService

1. SERVICE LAYER (service/)
   ├── base_service.go          → Common dependencies (Store, Logger, Validator)
   ├── user_service.go          → Handles user operations
   ├── product_service.go       → Uses UserService internally
   └── factory.go               → Creates and caches all services

2. HANDLER LAYER (server/http/handler/)
   ├── base_handler.go          → Common dependencies
   ├── user_handler.go          → Handles user HTTP requests
   ├── product_handler.go       → Uses ProductService (which uses UserService)
   └── factory.go               → Creates and caches all handlers

DEPENDENCY FLOW:
================

http.go (registerRoutes)
    ↓
GetProductHandler from HandlerFactory
    ↓
HandlerFactory creates ProductHandler with:
    - BaseHandler
    - ProductService
    - UserService
    ↓
ProductService has:
    - BaseService (Store, Logger, Validator)
    - UserService

STEP-BY-STEP EXAMPLE:
====================

1. In ServiceFactory.GetProductService():
   
   func (sf *ServiceFactory) GetProductService() *ProductService {
       if sf.productService != nil {
           return sf.productService
       }
       
       sf.productService = NewProductService(
           sf.store,
           sf.logger,
           sf.validator,
           sf.GetUserService(),  // Pass UserService as dependency
       )
       return sf.productService
   }

2. In ProductService struct:

   type ProductService struct {
       *BaseService              // Inherited: Store, Logger, Validator
       UserService *UserService  // Dependency on UserService
   }

3. In ProductService method (using UserService):

   func (ps *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
       // Validate product
       if err := ps.Validator.Struct(product); err != nil {
           ps.Logger.Errorw("validation failed", "error", err)
           return err
       }

       // Use UserService to validate the owner
       user, err := ps.UserService.GetUser(ctx, product.OwnerID)
       if err != nil {
           ps.Logger.Errorw("failed to get product owner", "error", err)
           return err
       }

       // Now proceed with product creation
       return nil // ps.Store.CreateProduct(ctx, product)
   }

4. In HandlerFactory.GetProductHandler():

   func (hf *HandlerFactory) GetProductHandler() *ProductHandler {
       if hf.productHandler != nil {
           return hf.productHandler
       }

       hf.productHandler = NewProductHandler(
           hf.baseHandler,
           hf.serviceFactory.GetProductService(),  // Includes UserService
           hf.serviceFactory.GetUserService(),     // Pass UserService too if needed
       )
       return hf.productHandler
   }

5. In http.go registerRoutes():

   protected := s.router.Group("/api")
   protected.Use(AuthMiddleware())
   {
       productHandler := s.HandlerFactory.GetProductHandler()
       products := protected.Group("/products")
       {
           products.POST("", productHandler.CreateProduct)
       }
   }

KEY POINTS:
===========

✅ Services are created ONCE and CACHED by the factory
✅ Services can depend on other services
✅ All dependencies flow through the factory
✅ The factory ensures singleton pattern (one instance per service)
✅ Easy to test: you can mock services at the factory level
✅ No circular dependencies: flow is always one direction

ADDING A NEW SERVICE WITH MULTIPLE DEPENDENCIES:
==================================================

1. Create the service (service/order_service.go):

   type OrderService struct {
       *BaseService
       ProductService *service.ProductService
       UserService    *service.UserService
   }

   func NewOrderService(
       store *store.Store,
       logger *zap.SugaredLogger,
       validator *validator.Validate,
       productService *service.ProductService,
       userService *service.UserService,
   ) *OrderService {
       return &OrderService{
           BaseService: NewBaseService(store, logger, validator),
           ProductService: productService,
           UserService: userService,
       }
   }

2. Add to ServiceFactory (service/factory.go):

   type ServiceFactory struct {
       // ... existing fields ...
       orderService *OrderService
   }

   func (sf *ServiceFactory) GetOrderService() *OrderService {
       if sf.orderService != nil {
           return sf.orderService
       }
       sf.orderService = NewOrderService(
           sf.store,
           sf.logger,
           sf.validator,
           sf.GetProductService(),
           sf.GetUserService(),
       )
       return sf.orderService
   }

3. Create handler (server/http/handler/order_handler.go)
4. Add to HandlerFactory
5. Register in http.go

That's it! Follow the same pattern for any new service.
*/
