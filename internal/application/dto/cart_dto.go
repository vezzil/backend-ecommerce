package dto

// CartCreateRequest represents the data needed to create a new cart
type CartCreateRequest struct {
	UserID     *string `json:"user_id,omitempty" validate:"omitempty,uuid4"`
	GuestToken string  `json:"guest_token,omitempty"`
}

// CartItemAddRequest represents the data needed to add an item to a cart
type CartItemAddRequest struct {
	ProductID string `json:"product_id" validate:"required,uuid4"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

// CartItemUpdateRequest represents the data needed to update a cart item
type CartItemUpdateRequest struct {
	Quantity int `json:"quantity" validate:"required,min=0"`
}

// CartResponse represents the cart data returned to the client
type CartResponse struct {
	ID         string             `json:"id"`
	UserID     *string            `json:"user_id,omitempty"`
	GuestToken string             `json:"guest_token,omitempty"`
	Items      []CartItemResponse `json:"items,omitempty"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
}

// CartItemResponse represents a cart item in the response
type CartItemResponse struct {
	ID        string           `json:"id"`
	ProductID string           `json:"product_id"`
	Product   *ProductResponse `json:"product,omitempty"`
	Quantity  int              `json:"quantity"`
	UnitPrice float64          `json:"unit_price"`
	CreatedAt string           `json:"created_at"`
}
