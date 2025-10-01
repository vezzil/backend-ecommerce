package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend-ecommerce/internal/application/dto"
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/service"
	"backend-ecommerce/pkg/response"
)

// ProductController handles product-related HTTP requests
type ProductController struct {
	productService service.ProductService
}

// NewProductController creates a new product controller
func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// RegisterRoutes registers product routes
func (pc *ProductController) RegisterRoutes(router *gin.RouterGroup) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("", pc.GetProducts)
		productGroup.POST("", pc.CreateProduct)
		productGroup.GET("/:id", pc.GetProduct)
		productGroup.PUT("/:id", pc.UpdateProduct)
		productGroup.DELETE("/:id", pc.DeleteProduct)
	}
}

// GetProducts handles GET /api/products
func (pc *ProductController) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryID := c.Query("category_id")

	var categoryIDPtr *string
	if categoryID != "" {
		categoryIDPtr = &categoryID
	}

	products, total, err := pc.productService.GetAllProducts(page, pageSize, categoryIDPtr)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch products: "+err.Error())
		return
	}

	response.Success(c, dto.PaginatedResponse{
		Success:  true,
		Data:     products,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

// GetProduct handles GET /api/products/:id
func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Product")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch product: "+err.Error())
		}
		return
	}

	response.Success(c, product)
}

// CreateProduct handles POST /api/products
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req dto.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	product := entity.Product{
		SKU:         req.SKU,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
		CategoryID:  req.CategoryID,
		IsActive:    req.IsActive,
	}

	created, err := pc.productService.CreateProduct(product)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create product: "+err.Error())
		return
	}

	response.Created(c, created)
}

// UpdateProduct handles PUT /api/products/:id
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	var req dto.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	// Get existing product
	existing, err := pc.productService.GetProductByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Product")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch product: "+err.Error())
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
	if req.Price != nil {
		existing.Price = *req.Price
	}
	if req.Currency != nil {
		existing.Currency = *req.Currency
	}
	if req.CategoryID != nil {
		existing.CategoryID = req.CategoryID
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	updated, err := pc.productService.UpdateProduct(*existing)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product: "+err.Error())
		return
	}

	response.Success(c, updated)
}

// DeleteProduct handles DELETE /api/products/:id
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	err := pc.productService.DeleteProduct(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Product")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to delete product: "+err.Error())
		}
		return
	}

	response.Success(c, nil)
}
