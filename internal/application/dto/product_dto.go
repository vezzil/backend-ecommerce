package dto

// ProductCreateRequest represents the data needed to create a new product
type ProductCreateRequest struct {
	SKU         string   `json:"sku,omitempty" validate:"omitempty,max=100"`
	Name        string   `json:"name" validate:"required,min=2,max=255"`
	Slug        string   `json:"slug,omitempty" validate:"omitempty,slug,max=255"`
	Description string   `json:"description,omitempty"`
	Price       float64  `json:"price" validate:"required,gt=0"`
	Currency    string   `json:"currency,omitempty" validate:"omitempty,iso4217"`
	CategoryID  *string  `json:"category_id,omitempty" validate:"omitempty,uuid4"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

// ProductUpdateRequest represents the data needed to update an existing product
type ProductUpdateRequest struct {
	SKU         *string  `json:"sku,omitempty" validate:"omitempty,max=100"`
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Slug        *string  `json:"slug,omitempty" validate:"omitempty,slug,max=255"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Currency    *string  `json:"currency,omitempty" validate:"omitempty,iso4217"`
	CategoryID  *string  `json:"category_id,omitempty" validate:"omitempty,uuid4"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

// ProductResponse represents the product data returned to the client
type ProductResponse struct {
	ID           string             `json:"id"`
	SKU          string             `json:"sku,omitempty"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Description  string             `json:"description,omitempty"`
	Price        float64            `json:"price"`
	Currency     string             `json:"currency"`
	CategoryID   *string            `json:"category_id,omitempty"`
	IsActive     bool               `json:"is_active"`
	Category     *CategoryResponse  `json:"category,omitempty"`
	Images       []ProductImageResponse `json:"images,omitempty"`
	Inventory    *InventoryResponse `json:"inventory,omitempty"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
}
