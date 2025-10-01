package dto

// CategoryCreateRequest represents the data needed to create a new category
type CategoryCreateRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=150"`
	Slug        string `json:"slug,omitempty" validate:"omitempty,slug,max=150"`
	Description string `json:"description,omitempty"`
}

// CategoryUpdateRequest represents the data needed to update an existing category
type CategoryUpdateRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=150"`
	Slug        *string `json:"slug,omitempty" validate:"omitempty,slug,max=150"`
	Description *string `json:"description,omitempty"`
}

// CategoryResponse represents the category data returned to the client
type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
