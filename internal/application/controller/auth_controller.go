package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"backend-ecommerce/internal/application/dto"
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/infrastructure/config"
)

// AuthController handles authentication related requests
type AuthController struct {
	db *gorm.DB
}

// NewAuthController creates a new AuthController
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db: db}
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterRequest true "Registration details"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser entity.User
	if err := ac.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Check if username already exists
	if err := ac.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		FullName: req.FullName,
		IsActive: true,
	}

	if err := ac.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate token pair
	tokenPair, err := config.GenerateTokenPair(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Set secure, HTTP-only cookie for refresh token
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		tokenPair.RefreshToken,
		int((7 * 24 * time.Hour).Seconds()), // 7 days
		"/api/auth/refresh",
		"", // domain
		true, // secure (set to true in production with HTTPS)
		true, // httpOnly
	)

	// Return user info and access token in response body
	c.JSON(http.StatusCreated, dto.AuthResponse{
		UserID:    user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FullName:  user.FullName,
		Token:     tokenPair.AccessToken,
		ExpiresAt: tokenPair.ExpiresAt.Format(time.RFC3339),
	})
}

// Login handles user login
// @Summary User login
// @Description Authenticate a user and return a JWT token pair
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user entity.User
	if err := ac.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is deactivated"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token pair
	tokenPair, err := config.GenerateTokenPair(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Set secure, HTTP-only cookie for refresh token
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		tokenPair.RefreshToken,
		int((7 * 24 * time.Hour).Seconds()), // 7 days
		"/api/auth/refresh",
		"", // domain
		true, // secure (set to true in production with HTTPS)
		true, // httpOnly
	)

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	ac.db.Save(&user)

	// Return user info and access token in response body
	c.JSON(http.StatusOK, dto.AuthResponse{
		UserID:    user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FullName:  user.FullName,
		Token:     tokenPair.AccessToken,
		ExpiresAt: tokenPair.ExpiresAt.Format(time.RFC3339),
	})
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Description Refresh an expired access token using a refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/refresh [post]
func (ac *AuthController) RefreshToken(c *gin.Context) {
	// Get refresh token from cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	// Get user ID from context (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Verify refresh token and get new token pair
	tokenPair, err := config.VerifyRefreshToken(userID.(uint), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Set new refresh token in cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		tokenPair.RefreshToken,
		int((7 * 24 * time.Hour).Seconds()),
		"/api/auth/refresh",
		"",
		true,  // secure
		true,  // httpOnly
	)

	// Return new access token
	c.JSON(http.StatusOK, dto.AuthResponse{
		Token:     tokenPair.AccessToken,
		ExpiresAt: tokenPair.ExpiresAt.Format(time.RFC3339),
	})
}

// Logout handles user logout
// @Summary User logout
// @Description Invalidate the current user's refresh token
// @Tags auth
// @Security Bearer
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/logout [post]
func (ac *AuthController) Logout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// Invalidate refresh token
	config.InvalidateRefreshToken(userID.(uint))

	// Clear refresh token cookie
	c.SetCookie(
		"refresh_token",
		"",
		-1, // Expire immediately
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// GetMe returns the current user's profile
// @Summary Get current user profile
// @Description Get the profile of the currently authenticated user
// @Tags auth
// @Security Bearer
// @Produce json
// @Success 200 {object} dto.AuthResponse
// @Failure 401 {object} map[string]string
// @Router /api/auth/me [get]
func (ac *AuthController) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user entity.User
	if err := ac.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
	})
}
