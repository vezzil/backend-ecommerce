package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"backend-ecommerce/internal/auth/jwtmanager"
)

// JWTConfig holds the JWT configuration
var JWT *jwtmanager.Manager

// InitJWT initializes the JWT manager with configuration
func InitJWT() {
	cfg := Get()
	JWT = jwtmanager.New(
		cfg.JWT.Secret,
		cfg.JWT.Issuer,
		cfg.JWT.ExpireIn,
	)
}

// TokenResponse represents the JWT token response
type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// GenerateToken generates a new JWT token for a user
func GenerateToken(userID uint) (*TokenResponse, error) {
	tokenString, err := JWT.Sign(userID)
	if err != nil {
		return nil, err
	}

	// Parse the token to get expiration time
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwtmanager.Claims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtmanager.Claims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt.Time,
	}, nil
}
