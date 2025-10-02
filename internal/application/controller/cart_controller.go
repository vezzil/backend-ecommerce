package controller

import (
	"backend-ecommerce/internal/application/service"
	"backend-ecommerce/internal/application/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CartController handles cart-related HTTP requests
type CartController struct{

}

// GetOrCreateCart handles GET /api/carts
// @Summary Get or create user's shopping cart
// @Description Retrieves the user's shopping cart or creates a new one if it doesn't exist
// @Tags Cart
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} dto.ResponseDto "Successfully retrieved or created cart"
// @Failure 401 {object} dto.ResponseDto "Unauthorized"
// @Failure 500 {object} dto.ResponseDto "Internal server error"
// @Router /api/carts [get]
func (cc *CartController) GetOrCreateCart(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, dto.Fail("User not authenticated"))
        return
    }

    response := service.ICartService.GetOrCreateCart(&userID, "")
    c.JSON(http.StatusOK, response)
}

// AddCartItem handles POST /api/carts/items
// @Summary Add item to cart
// @Description Adds a product to the user's shopping cart
// @Tags Cart
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.AddCartItemRequest true "Add item request"
// @Success 200 {object} dto.ResponseDto "Item added successfully"
// @Failure 400 {object} dto.ResponseDto "Invalid request data"
// @Failure 401 {object} dto.ResponseDto "Unauthorized"
// @Router /api/carts/items [post]
func (cc *CartController) AddCartItem(c *gin.Context) {

}

// UpdateCartItem handles PUT /api/carts/items/:item_id
// @Summary Update cart item quantity
// @Description Updates the quantity of an item in the cart
// @Tags Cart
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param item_id path string true "Cart Item ID"
// @Param request body dto.UpdateCartItemRequest true "Update item request"
// @Success 200 {object} dto.ResponseDto "Item updated successfully"
// @Failure 400 {object} dto.ResponseDto "Invalid request data"
// @Failure 401 {object} dto.ResponseDto "Unauthorized"
// @Router /api/carts/items/{item_id} [put]
func (cc *CartController) UpdateCartItem(c *gin.Context) {

}

// RemoveCartItem handles DELETE /api/carts/items/:item_id
// @Summary Remove item from cart
// @Description Removes an item from the shopping cart
// @Tags Cart
// @Security ApiKeyAuth
// @Param item_id path string true "Cart Item ID"
// @Success 200 {object} dto.ResponseDto "Item removed successfully"
// @Failure 400 {object} dto.ResponseDto "Invalid item ID"
// @Failure 401 {object} dto.ResponseDto "Unauthorized"
// @Router /api/carts/items/{item_id} [delete]
func (cc *CartController) RemoveCartItem(c *gin.Context) {

}

// ClearCart handles DELETE /api/carts
// @Summary Clear shopping cart
// @Description Removes all items from the shopping cart
// @Tags Cart
// @Security ApiKeyAuth
// @Success 200 {object} dto.ResponseDto "Cart cleared successfully"
// @Failure 401 {object} dto.ResponseDto "Unauthorized"
// @Router /api/carts [delete]
func (cc *CartController) ClearCart(c *gin.Context) {

}
