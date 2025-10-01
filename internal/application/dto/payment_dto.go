package dto

// PaymentCreateRequest represents the data needed to process a payment
type PaymentCreateRequest struct {
	OrderID     string  `json:"order_id" validate:"required,uuid4"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Currency    string  `json:"currency" validate:"required,iso4217"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	CardToken   string  `json:"card_token,omitempty"`
}

// PaymentResponse represents the payment data returned to the client
type PaymentResponse struct {
	ID                string            `json:"id"`
	OrderID           string            `json:"order_id"`
	Provider          string            `json:"provider"`
	ProviderPaymentID string            `json:"provider_payment_id,omitempty"`
	Amount            float64           `json:"amount"`
	Currency          string            `json:"currency"`
	Status            string            `json:"status"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
}
