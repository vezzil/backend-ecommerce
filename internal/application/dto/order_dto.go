package dto

// OrderCreateRequest represents the data needed to create a new order
type OrderCreateRequest struct {
	UserID            string   `json:"user_id" validate:"required,uuid4"`
	ShippingAddressID string   `json:"shipping_address_id" validate:"required,uuid4"`
	BillingAddressID  string   `json:"billing_address_id" validate:"required,uuid4"`
	CartID            string   `json:"cart_id" validate:"required,uuid4"`
}

// OrderUpdateRequest represents the data needed to update an order
type OrderUpdateRequest struct {
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=pending paid processing shipped cancelled completed"`
}

// OrderResponse represents the order data returned to the client
type OrderResponse struct {
	ID                string             `json:"id"`
	UserID            string             `json:"user_id"`
	Status            string             `json:"status"`
	TotalAmount       float64            `json:"total_amount"`
	Currency          string             `json:"currency"`
	ShippingAddressID string             `json:"shipping_address_id"`
	BillingAddressID  string             `json:"billing_address_id"`
	Items             []OrderItemResponse `json:"items,omitempty"`
	Payment           *PaymentResponse    `json:"payment,omitempty"`
	CreatedAt         string             `json:"created_at"`
	UpdatedAt         string             `json:"updated_at"`
}

// OrderItemResponse represents an order item in the response
type OrderItemResponse struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	SKU         string  `json:"sku,omitempty"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}
