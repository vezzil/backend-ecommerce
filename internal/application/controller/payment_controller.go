package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/service"
	"backend-ecommerce/pkg/response"
)

// PaymentController handles payment-related HTTP requests
type PaymentController struct {
	paymentService service.PaymentService
}

// NewPaymentController creates a new payment controller
func NewPaymentController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

// RegisterRoutes registers payment routes
func (pc *PaymentController) RegisterRoutes(router *gin.RouterGroup) {
	paymentGroup := router.Group("/payments")
	{
		paymentGroup.POST("", pc.ProcessPayment)
		paymentGroup.GET("/:id", pc.GetPayment)
		paymentGroup.GET("", pc.GetPaymentsByOrder)
		paymentGroup.POST("/:id/refund", pc.RefundPayment)
	}
}

// ProcessPayment handles POST /api/payments
func (pc *PaymentController) ProcessPayment(c *gin.Context) {
	var req struct {
		OrderID       string  `json:"order_id" binding:"required"`
		Amount        float64 `json:"amount" binding:"required,gt=0"`
		PaymentMethod string  `json:"payment_method" binding:"required"`
		CardToken     string  `json:"card_token,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	payment := entity.Payment{
		OrderID:     req.OrderID,
		Amount:      req.Amount,
		Provider:    req.PaymentMethod,
		Status:      "pending",
	}

	processed, err := pc.paymentService.ProcessPayment(payment)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to process payment: "+err.Error())
		return
	}

	response.Created(c, processed)
}

// GetPayment handles GET /api/payments/:id
func (pc *PaymentController) GetPayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Payment ID is required")
		return
	}

	payment, err := pc.paymentService.GetPaymentByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Payment")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch payment: "+err.Error())
		}
		return
	}

	response.Success(c, payment)
}

// GetPaymentsByOrder handles GET /api/payments
func (pc *PaymentController) GetPaymentsByOrder(c *gin.Context) {
	orderID := c.Query("order_id")
	if orderID == "" {
		response.Error(c, http.StatusBadRequest, "Order ID is required")
		return
	}

	payments, err := pc.paymentService.GetPaymentsByOrderID(orderID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch payments: "+err.Error())
		return
	}

	response.Success(c, payments)
}

// RefundPayment handles POST /api/payments/:id/refund
func (pc *PaymentController) RefundPayment(c *gin.Context) {
	paymentID := c.Param("id")
	if paymentID == "" {
		response.Error(c, http.StatusBadRequest, "Payment ID is required")
		return
	}

	var req struct {
		Amount float64 `json:"amount,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	err := pc.paymentService.RefundPayment(paymentID, req.Amount)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Payment")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to process refund: "+err.Error())
		}
		return
	}

	response.Success(c, nil)
}
