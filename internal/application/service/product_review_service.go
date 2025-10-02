package service

import (
	"backend-ecommerce/internal/application/entity"
)

type productReviewService struct {
}

func (s *productReviewService) CreateReview(review entity.ProductReview) (*entity.ProductReview, error) {
	// Add any validation or business logic here
	return nil, nil
}

func (s *productReviewService) GetProductReviews(productID string, page, pageSize int) ([]entity.ProductReview, int64, error) {
	return []entity.ProductReview{}, 0, nil
}

func (s *productReviewService) GetReviewByID(id string) (*entity.ProductReview, error) {
	return nil, nil
}

func (s *productReviewService) UpdateReview(review entity.ProductReview) (*entity.ProductReview, error) {
	// Add any validation or business logic here
	return nil, nil
}

func (s *productReviewService) DeleteReview(id string) error {
	return nil
}
