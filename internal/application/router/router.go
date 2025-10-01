package router

import (
	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/controller"
	"backend-ecommerce/internal/application/service"
)

// Register registers all HTTP routes on the given engine.
func Register(r *gin.Engine) {
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
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)
	orderController := controller.NewOrderController(orderService)
	paymentController := controller.NewPaymentController(paymentService)
	productReviewController := controller.NewProductReviewController(productReviewService)

	// API routes
	api := r.Group("/api")
	{
		// User routes
		userController.RegisterRoutes(api)

		// Category routes
		categoryController.RegisterRoutes(api)

		// Product routes
		productController.RegisterRoutes(api)

		// Cart routes
		cartController.RegisterRoutes(api)

		// Order routes
		orderController.RegisterRoutes(api)

		// Payment routes
		paymentController.RegisterRoutes(api)

		// Product review routes
		productReviewController.RegisterRoutes(api)

		// Health check.
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}
}
