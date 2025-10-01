package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type productService struct {
	repo repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetAllProducts(page, pageSize int, categoryID *string) ([]entity.Product, int64, error) {
	return s.repo.FindAll(page, pageSize, categoryID)
}

func (s *productService) GetProductByID(id string) (*entity.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) CreateProduct(product entity.Product) (*entity.Product, error) {
	// Add any business logic here before creating
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(product entity.Product) (*entity.Product, error) {
	// Add any business logic here before updating
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}
