package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"backend-ecommerce/internal/infrastructure/config"
)

// AuthMiddleware is a middleware that checks for a valid JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Extract the token from the header
		// Format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]

		// Verify the token
		claims, err := config.JWT.Verify(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Add user ID to context
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

// GetUserIDFromContext retrieves the user ID from the context
func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	// Convert to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		return 0, false
	}

	return userIDUint, true
}

// AdminMiddleware checks if the user is an admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := GetUserIDFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// TODO: Implement admin check logic here
		// For now, we'll just allow any authenticated user
		// In a real app, you would check the user's role here

		c.Next()
	}
}
