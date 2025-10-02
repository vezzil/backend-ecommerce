package service

import (
	"backend-ecommerce/internal/application/entity"
)

type categoryService struct {
}



func (s *categoryService) GetAllCategories(page, pageSize int) ([]entity.Category, int64, error) {
	return []entity.Category{}, 0, nil
}

func (s *categoryService) GetCategoryByID(id string) (*entity.Category, error) {
	return nil, nil
}

func (s *categoryService) CreateCategory(category entity.Category) (*entity.Category, error) {
	// Add any business logic here before creating
	return nil, nil
}

func (s *categoryService) UpdateCategory(category entity.Category) (*entity.Category, error) {
	// Add any business logic here before updating
	return nil, nil
}

func (s *categoryService) DeleteCategory(id string) error {
	return nil
}
