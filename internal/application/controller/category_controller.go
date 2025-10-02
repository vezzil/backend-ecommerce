package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/dto"
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/service"
)

// CategoryController handles category-related HTTP requests
type CategoryController struct {
	categoryService service.CategoryService
}

// NewCategoryController creates a new category controller
func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// RegisterRoutes registers category routes
func (cc *CategoryController) RegisterRoutes(router *gin.RouterGroup) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("", cc.GetCategories)
		categoryGroup.POST("", cc.CreateCategory)
		categoryGroup.GET("/:id", cc.GetCategory)
		categoryGroup.PUT("/:id", cc.UpdateCategory)
		categoryGroup.DELETE("/:id", cc.DeleteCategory)
	}
}

// GetCategories handles GET /api/categories
func (cc *CategoryController) GetCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	categories, total, err := cc.categoryService.GetAllCategories(page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
		return
	}

	response.Success(c, dto.PaginatedResponse{
		Success:  true,
		Data:     categories,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

// GetCategory handles GET /api/categories/:id
func (cc *CategoryController) GetCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Category ID is required")
		return
	}

	category, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Category")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch category: "+err.Error())
		}
		return
	}

	response.Success(c, category)
}

// CreateCategory handles POST /api/categories
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var req dto.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	category := entity.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	created, err := cc.categoryService.CreateCategory(category)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create category: "+err.Error())
		return
	}

	response.Created(c, created)
}

// UpdateCategory handles PUT /api/categories/:id
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Category ID is required")
		return
	}

	var req dto.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	// Get existing category
	existing, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Category")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch category: "+err.Error())
		}
		return
	}

	// Update fields if provided
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Slug != nil {
		existing.Slug = *req.Slug
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}

	updated, err := cc.categoryService.UpdateCategory(*existing)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update category: "+err.Error())
		return
	}

	response.Success(c, updated)
}

// DeleteCategory handles DELETE /api/categories/:id
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Category ID is required")
		return
	}

	err := cc.categoryService.DeleteCategory(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Category")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to delete category: "+err.Error())
		}
		return
	}

	response.Success(c, nil)
}
