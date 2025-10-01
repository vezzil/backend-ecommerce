package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/service"
	"backend-ecommerce/pkg/response"
)

// OrderController handles order-related HTTP requests
type OrderController struct {
	orderService service.OrderService
}

// NewOrderController creates a new order controller
func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

// RegisterRoutes registers order routes
func (oc *OrderController) RegisterRoutes(router *gin.RouterGroup) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", oc.CreateOrder)
		orderGroup.GET("", oc.GetUserOrders)
		orderGroup.GET("/:id", oc.GetOrder)
		orderGroup.PUT("/:id/cancel", oc.CancelOrder)
	}
}

// CreateOrder handles POST /api/orders
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var req struct {
		CartID            string `json:"cart_id" binding:"required"`
		ShippingAddressID string `json:"shipping_address_id" binding:"required"`
		BillingAddressID  string `json:"billing_address_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	// In a real app, you'd get the user ID from the auth context
	userID := "" // c.GetString("user_id") // Uncomment when auth is implemented

	order := entity.Order{
		UserID:            userID,
		ShippingAddressID: req.ShippingAddressID,
		BillingAddressID:  req.BillingAddressID,
	}

	created, err := oc.orderService.CreateOrder(order)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create order: "+err.Error())
		return
	}

	response.Created(c, created)
}

// GetOrder handles GET /api/orders/:id
func (oc *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Order ID is required")
		return
	}

	// In a real app, you'd verify the user has permission to view this order
	// userID := c.GetString("user_id")

	order, err := oc.orderService.GetOrderByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Order")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch order: "+err.Error())
		}
		return
	}

	// In a real app, you'd verify the user has permission to view this order
	// if order.UserID != userID {
	// 	response.Forbidden(c, "You don't have permission to view this order")
	// 	return
	// }

	response.Success(c, order)
}

// GetUserOrders handles GET /api/orders
func (oc *OrderController) GetUserOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// In a real app, you'd get the user ID from the auth context
	userID := "" // c.GetString("user_id") // Uncomment when auth is implemented

	orders, total, err := oc.orderService.GetUserOrders(userID, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch orders: "+err.Error())
		return
	}

	response.Success(c, gin.H{
		"data":      orders,
		"page":      page,
		"page_size": pageSize,
		"total":     total,
	})
}

// CancelOrder handles PUT /api/orders/:id/cancel
func (oc *OrderController) CancelOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Order ID is required")
		return
	}

	// In a real app, you'd verify the user has permission to cancel this order
	// userID := c.GetString("user_id")

	err := oc.orderService.CancelOrder(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Order")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to cancel order: "+err.Error())
		}
		return
	}

	response.Success(c, nil)
}
