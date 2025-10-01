package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type productReviewService struct {
	repo repository.ProductReviewRepository
}

// NewProductReviewService creates a new product review service
func NewProductReviewService(repo repository.ProductReviewRepository) ProductReviewService {
	return &productReviewService{
		repo: repo,
	}
}

func (s *productReviewService) CreateReview(review entity.ProductReview) (*entity.ProductReview, error) {
	// Add any validation or business logic here
	return s.repo.Create(review)
}

func (s *productReviewService) GetProductReviews(productID string, page, pageSize int) ([]entity.ProductReview, int64, error) {
	return s.repo.FindByProductID(productID, page, pageSize)
}

func (s *productReviewService) GetReviewByID(id string) (*entity.ProductReview, error) {
	return s.repo.FindByID(id)
}

func (s *productReviewService) UpdateReview(review entity.ProductReview) (*entity.ProductReview, error) {
	// Add any validation or business logic here
	return s.repo.Update(review)
}

func (s *productReviewService) DeleteReview(id string) error {
	return s.repo.Delete(id)
}
