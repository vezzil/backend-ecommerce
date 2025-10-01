package dto

// ProductImageCreateRequest represents the data needed to add an image to a product
type ProductImageCreateRequest struct {
	ProductID string `json:"product_id" validate:"required,uuid4"`
	URL       string `json:"url" validate:"required,url,max=1000"`
	AltText   string `json:"alt_text,omitempty" validate:"max=255"`
	SortOrder int    `json:"sort_order,omitempty"`
}

// ProductImageUpdateRequest represents the data needed to update a product image
type ProductImageUpdateRequest struct {
	URL       *string `json:"url,omitempty" validate:"omitempty,url,max=1000"`
	AltText   *string `json:"alt_text,omitempty" validate:"max=255"`
	SortOrder *int    `json:"sort_order,omitempty"`
}

// ProductImageResponse represents the product image data returned to the client
type ProductImageResponse struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	URL       string `json:"url"`
	AltText   string `json:"alt_text,omitempty"`
	SortOrder int    `json:"sort_order"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
