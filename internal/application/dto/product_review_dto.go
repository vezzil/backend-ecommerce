package dto

// ProductReviewCreateRequest represents the data needed to create a product review
type ProductReviewCreateRequest struct {
	ProductID string  `json:"product_id" validate:"required,uuid4"`
	Rating    int     `json:"rating" validate:"required,min=1,max=5"`
	Title     string  `json:"title" validate:"required,min=2,max=200"`
	Comment   string  `json:"comment" validate:"required,min=10"`
}

// ProductReviewUpdateRequest represents the data needed to update a product review
type ProductReviewUpdateRequest struct {
	Rating  *int    `json:"rating,omitempty" validate:"omitempty,min=1,max=5"`
	Title   *string `json:"title,omitempty" validate:"omitempty,min=2,max=200"`
	Comment *string `json:"comment,omitempty" validate:"omitempty,min=10"`
}

// ProductReviewResponse represents the product review data returned to the client
type ProductReviewResponse struct {
	ID          string          `json:"id"`
	ProductID   string          `json:"product_id"`
	UserID      string          `json:"user_id"`
	User        *UserResponse   `json:"user,omitempty"`
	Rating      int             `json:"rating"`
	Title       string          `json:"title"`
	Comment     string          `json:"comment"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}
