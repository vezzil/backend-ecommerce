package service

import (
	"backend-ecommerce/internal/application/entity"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	// Get all users with pagination
	GetAllUsers(page, pageSize int) ([]entity.User, int64, error)
	// Get user by ID
	GetUserByID(id string) (*entity.User, error)
	// Create a new user
	CreateUser(user entity.User) (*entity.User, error)
	// Update an existing user
	UpdateUser(user entity.User) (*entity.User, error)
	// Delete a user
	DeleteUser(id string) error
}

// CategoryService defines the interface for category-related business logic
type CategoryService interface {
	// GetAllCategories retrieves all categories with pagination support
	// Returns a list of categories, total count, and any error encountered
	GetAllCategories(page, pageSize int) ([]entity.Category, int64, error)
	
	// GetCategoryByID retrieves a single category by its ID
	// Returns the category if found, nil if not found, and any error encountered
	GetCategoryByID(id string) (*entity.Category, error)
	
	// CreateCategory creates a new category
	// Returns the created category and any error encountered
	CreateCategory(category entity.Category) (*entity.Category, error)
	
	// UpdateCategory updates an existing category
	// Returns the updated category and any error encountered
	UpdateCategory(category entity.Category) (*entity.Category, error)
	
	// DeleteCategory removes a category by its ID
	// Returns an error if the operation fails
	DeleteCategory(id string) error
}

// ProductService defines the interface for product-related business logic
type ProductService interface {
	// GetAllProducts retrieves all products with optional category filtering and pagination
	// Returns a list of products, total count, and any error encountered
	GetAllProducts(page, pageSize int, categoryID *string) ([]entity.Product, int64, error)
	
	// GetProductByID retrieves a single product by its ID
	// Returns the product if found, nil if not found, and any error encountered
	GetProductByID(id string) (*entity.Product, error)
	
	// CreateProduct creates a new product
	// Returns the created product and any error encountered
	CreateProduct(product entity.Product) (*entity.Product, error)
	
	// UpdateProduct updates an existing product
	// Returns the updated product and any error encountered
	UpdateProduct(product entity.Product) (*entity.Product, error)
	
	// DeleteProduct removes a product by its ID
	// Returns an error if the operation fails
	DeleteProduct(id string) error
}

// CartService defines the interface for shopping cart business logic
type CartService interface {
	// GetUserCart retrieves the shopping cart for a specific user
	// Returns the user's cart and any error encountered
	GetUserCart(userID string) (*entity.Cart, error)
	
	// AddToCart adds an item to the user's shopping cart
	// Returns the updated cart and any error encountered
	AddToCart(userID string, item entity.CartItem) (*entity.Cart, error)
	
	// UpdateCartItem updates the quantity of a specific item in the cart
	// Returns the updated cart and any error encountered
	UpdateCartItem(userID, itemID string, quantity int) (*entity.Cart, error)
	
	// RemoveFromCart removes an item from the user's shopping cart
	// Returns an error if the operation fails
	RemoveFromCart(userID, itemID string) error
	
	// ClearCart removes all items from the user's shopping cart
	// Returns an error if the operation fails
	ClearCart(userID string) error
}

// OrderService defines the interface for order business logic
type OrderService interface {
	// CreateOrder creates a new order
	// Returns the created order and any error encountered
	CreateOrder(order entity.Order) (*entity.Order, error)
	
	// GetOrderByID retrieves a specific order by its ID
	// Returns the order if found, nil if not found, and any error encountered
	GetOrderByID(id string) (*entity.Order, error)
	
	// GetUserOrders retrieves all orders for a specific user with pagination
	// Returns a list of orders, total count, and any error encountered
	GetUserOrders(userID string, page, pageSize int) ([]entity.Order, int64, error)
	
	// UpdateOrderStatus updates the status of an existing order
	// Returns an error if the operation fails
	UpdateOrderStatus(id string, status entity.OrderStatus) error
	
	// CancelOrder cancels an existing order
	// Returns an error if the operation fails
	CancelOrder(id string) error
}

// PaymentService defines the interface for payment processing logic
type PaymentService interface {
	// CreatePayment creates a new payment record
	// Returns the created payment and any error encountered
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	
	// GetPaymentByID retrieves a specific payment by its ID
	// Returns the payment if found, nil if not found, and any error encountered
	GetPaymentByID(id string) (*entity.Payment, error)
	
	// ProcessPayment processes a payment with the payment provider
	// Returns the updated payment and any error encountered
	ProcessPayment(paymentID string) (*entity.Payment, error)
	
	// RefundPayment processes a refund for a payment
	// Returns the updated payment and any error encountered
	RefundPayment(paymentID string, amount float64) (*entity.Payment, error)
	
	// GetOrderPayments retrieves all payments for a specific order
	// Returns a list of payments and any error encountered
	GetOrderPayments(orderID string) ([]entity.Payment, error)
}

// ProductReviewService defines the interface for product review business logic
type ProductReviewService interface {
	// CreateReview creates a new product review
	// Returns the created review and any error encountered
	CreateReview(review entity.ProductReview) (*entity.ProductReview, error)
	
	// GetReviewByID retrieves a specific review by its ID
	// Returns the review if found, nil if not found, and any error encountered
	GetReviewByID(id string) (*entity.ProductReview, error)
	
	// GetProductReviews retrieves all reviews for a specific product with pagination
	// Returns a list of reviews, total count, and any error encountered
	GetProductReviews(productID string, page, pageSize int) ([]entity.ProductReview, int64, error)
	
	// UpdateReview updates an existing product review
	// Returns the updated review and any error encountered
	UpdateReview(review entity.ProductReview) (*entity.ProductReview, error)
	
	// DeleteReview removes a review by its ID
	// Returns an error if the operation fails
	DeleteReview(id string) error
}
