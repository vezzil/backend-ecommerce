package dto

// AddressCreateRequest represents the data needed to create a new address
type AddressCreateRequest struct {
	UserID     string `json:"user_id" validate:"required,uuid"`
	Label      string `json:"label" validate:"required,max=100"`
	Street     string `json:"street" validate:"required,max=255"`
	City       string `json:"city" validate:"required,max=100"`
	State      string `json:"state" validate:"required,max=100"`
	PostalCode string `json:"postal_code" validate:"required,max=50"`
	Country    string `json:"country" validate:"required,max=100"`
	Phone      string `json:"phone" validate:"omitempty,max=50"`
}

// AddressUpdateRequest represents the data needed to update an existing address
type AddressUpdateRequest struct {
	Label      string `json:"label,omitempty" validate:"omitempty,max=100"`
	Street     string `json:"street,omitempty" validate:"omitempty,max=255"`
	City       string `json:"city,omitempty" validate:"omitempty,max=100"`
	State      string `json:"state,omitempty" validate:"omitempty,max=100"`
	PostalCode string `json:"postal_code,omitempty" validate:"omitempty,max=50"`
	Country    string `json:"country,omitempty" validate:"omitempty,max=100"`
	Phone      string `json:"phone,omitempty" validate:"omitempty,max=50"`
}

// AddressResponse represents the address data returned to the client
type AddressResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Label      string `json:"label"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Phone      string `json:"phone,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
