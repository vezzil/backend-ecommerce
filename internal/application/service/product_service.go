package service

import (
	"backend-ecommerce/internal/application/entity"
)

type productService struct {
}

func (s *productService) GetAllProducts(page, pageSize int, categoryID *string) ([]entity.Product, int64, error) {
	return []entity.Product{}, 0, nil
}

func (s *productService) GetProductByID(id string) (*entity.Product, error) {
	return nil, nil
}

func (s *productService) CreateProduct(product entity.Product) (*entity.Product, error) {
	return nil, nil
}

func (s *productService) UpdateProduct(product entity.Product) (*entity.Product, error) {
	return nil, nil
}

func (s *productService) DeleteProduct(id string) error {
	return nil
}
