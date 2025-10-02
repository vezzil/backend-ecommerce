package controller

import (
	"backend-ecommerce/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CartController handles cart-related HTTP requests
type CartController struct {
}

// GetOrCreateCart handles GET /api/carts
func (cc *CartController) GetOrCreateCart(c *gin.Context) {
	userID := c.GetString("user_id")
	
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	response := service.ICartService.GetOrCreateCart(userID, guestToken)

	c.JSON(http.StatusOK, response)
}

// AddCartItem handles POST /api/carts/items
func (cc *CartController) AddCartItem(c *gin.Context) {
	var req struct {
		CartID    string `json:"cart_id" binding:"required"`
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	item := entity.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	cart, err := cc.cartService.AddCartItem(req.CartID, item)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to add item to cart: "+err.Error())
		return
	}

	response.Success(c, cart)
}

// UpdateCartItem handles PUT /api/carts/items/:item_id
func (cc *CartController) UpdateCartItem(c *gin.Context) {
	cartID := c.Query("cart_id")
	if cartID == "" {
		response.Error(c, http.StatusBadRequest, "Cart ID is required")
		return
	}

	itemID := c.Param("item_id")
	if itemID == "" {
		response.Error(c, http.StatusBadRequest, "Item ID is required")
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required,min=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	cart, err := cc.cartService.UpdateCartItem(cartID, itemID, req.Quantity)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update cart item: "+err.Error())
		return
	}

	response.Success(c, cart)
}

// RemoveCartItem handles DELETE /api/carts/items/:item_id
func (cc *CartController) RemoveCartItem(c *gin.Context) {
	cartID := c.Query("cart_id")
	if cartID == "" {
		response.Error(c, http.StatusBadRequest, "Cart ID is required")
		return
	}

	itemID := c.Param("item_id")
	if itemID == "" {
		response.Error(c, http.StatusBadRequest, "Item ID is required")
		return
	}

	err := cc.cartService.RemoveCartItem(cartID, itemID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to remove item from cart: "+err.Error())
		return
	}

	response.Success(c, nil)
}

// ClearCart handles DELETE /api/carts
func (cc *CartController) ClearCart(c *gin.Context) {
	cartID := c.Query("cart_id")
	if cartID == "" {
		response.Error(c, http.StatusBadRequest, "Cart ID is required")
		return
	}

	err := cc.cartService.ClearCart(cartID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to clear cart: "+err.Error())
		return
	}

	response.Success(c, nil)
}
