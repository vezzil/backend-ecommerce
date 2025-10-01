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

// ProductReviewController handles product review-related HTTP requests
type ProductReviewController struct {
	reviewService service.ProductReviewService
}

// NewProductReviewController creates a new product review controller
func NewProductReviewController(reviewService service.ProductReviewService) *ProductReviewController {
	return &ProductReviewController{
		reviewService: reviewService,
	}
}

// RegisterRoutes registers product review routes
func (prc *ProductReviewController) RegisterRoutes(router *gin.RouterGroup) {
	reviewGroup := router.Group("/product-reviews")
	{
		reviewGroup.POST("", prc.CreateReview)
		reviewGroup.GET("", prc.GetProductReviews)
		reviewGroup.GET("/:id", prc.GetReview)
		reviewGroup.PUT("/:id", prc.UpdateReview)
		reviewGroup.DELETE("/:id", prc.DeleteReview)
	}
}

// CreateReview handles POST /api/product-reviews
func (prc *ProductReviewController) CreateReview(c *gin.Context) {
	var req dto.ProductReviewCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	// In a real app, you'd get the user ID from the auth context
	userID := "" // c.GetString("user_id") // Uncomment when auth is implemented

	review := entity.ProductReview{
		ProductID: req.ProductID,
		UserID:    userID,
		Rating:    req.Rating,
		Title:     req.Title,
		Comment:   req.Comment,
	}

	created, err := prc.reviewService.CreateReview(review)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create review: "+err.Error())
		return
	}

	response.Created(c, created)
}

// GetReview handles GET /api/product-reviews/:id
func (prc *ProductReviewController) GetReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Review ID is required")
		return
	}

	review, err := prc.reviewService.GetReviewByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Review")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch review: "+err.Error())
		}
		return
	}

	response.Success(c, review)
}

// GetProductReviews handles GET /api/product-reviews
func (prc *ProductReviewController) GetProductReviews(c *gin.Context) {
	productID := c.Query("product_id")
	if productID == "" {
		response.Error(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	reviews, total, err := prc.reviewService.GetProductReviews(productID, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch reviews: "+err.Error())
		return
	}

	response.Success(c, gin.H{
		"data":      reviews,
		"page":      page,
		"page_size": pageSize,
		"total":     total,
	})
}

// UpdateReview handles PUT /api/product-reviews/:id
func (prc *ProductReviewController) UpdateReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Review ID is required")
		return
	}

	// In a real app, you'd verify the user has permission to update this review
	// userID := c.GetString("user_id")

	var req dto.ProductReviewUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input data")
		return
	}

	// Get existing review
	existing, err := prc.reviewService.GetReviewByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Review")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch review: "+err.Error())
		}
		return
	}

	// In a real app, you'd verify the user has permission to update this review
	// if existing.UserID != userID {
	// 	response.Forbidden(c, "You don't have permission to update this review")
	// 	return
	// }

	// Update fields if provided
	if req.Rating != nil {
		existing.Rating = *req.Rating
	}
	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Comment != nil {
		existing.Comment = *req.Comment
	}

	updated, err := prc.reviewService.UpdateReview(*existing)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update review: "+err.Error())
		return
	}

	response.Success(c, updated)
}

// DeleteReview handles DELETE /api/product-reviews/:id
func (prc *ProductReviewController) DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Review ID is required")
		return
	}

	// In a real app, you'd verify the user has permission to delete this review
	// userID := c.GetString("user_id")

	// Get existing review (to verify ownership)
	existing, err := prc.reviewService.GetReviewByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			response.NotFound(c, "Review")
		} else {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch review: "+err.Error())
		}
		return
	}

	// In a real app, you'd verify the user has permission to delete this review
	// if existing.UserID != userID {
	// 	response.Forbidden(c, "You don't have permission to delete this review")
	// 	return
	// }

	err = prc.reviewService.DeleteReview(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete review: "+err.Error())
		return
	}

	response.Success(c, nil)
}
