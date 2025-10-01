package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type categoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService creates a new category service
func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) GetAllCategories(page, pageSize int) ([]entity.Category, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *categoryService) GetCategoryByID(id string) (*entity.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) CreateCategory(category entity.Category) (*entity.Category, error) {
	// Add any business logic here before creating
	return s.repo.Create(category)
}

func (s *categoryService) UpdateCategory(category entity.Category) (*entity.Category, error) {
	// Add any business logic here before updating
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
