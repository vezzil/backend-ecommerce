package router

import (
	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/controller"
	"backend-ecommerce/internal/application/service"
	"backend-ecommerce/internal/auth"
	"gorm.io/gorm"
)

// Register registers all HTTP routes on the given engine.
func Register(r *gin.Engine, db *gorm.DB) {
	// Initialize services
	userService := service.NewUserService()
	categoryService := service.NewCategoryService()
	productService := service.NewProductService()
	cartService := service.NewCartService()
	orderService := service.NewOrderService()
	paymentService := service.NewPaymentService()
	productReviewService := service.NewProductReviewService()

	// Initialize controllers
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(db)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)
	paymentController := controller.NewPaymentController(paymentService)
	productReviewController := controller.NewProductReviewController(productReviewService)

	// API routes
	api := r.Group("/api")

	// Public routes (no authentication required)
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// Auth endpoints
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", authController.Register)
			authGroup.POST("/login", authController.Login)
		}

		// Serve Swagger UI and docs
		// api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Protected API group (all routes under this group require authentication)
	protectedAPI := api.Group("")
	protectedAPI.Use(auth.AuthMiddleware())
	{
		// Auth protected routes
		authProtected := protectedAPI.Group("/auth")
		{
			authProtected.GET("/me", authController.GetMe)
		}

		// Register all other routes under the protected group
		userController.RegisterRoutes(protectedAPI)
		categoryController.RegisterRoutes(protectedAPI)
		productController.RegisterRoutes(protectedAPI)
		cartController.RegisterRoutes(protectedAPI)
		orderController.RegisterRoutes(protectedAPI)
		paymentController.RegisterRoutes(protectedAPI)
		productReviewController.RegisterRoutes(protectedAPI)
	}
