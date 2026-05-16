package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/regalangcom/go-shop-api/internal/config"
)

// claims struct for JWT payload
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// generate access token and refresh token
func GenerateTokenPair(cfg *config.JwtConfig, userID uint, email string, role string) (accessToken, refreshToken string, err error) {
	// access token
	accessClaims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JwtExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := at.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return "", "", err
	}

	// refresh token
	refreshTokenClaims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.RefreshToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := rt.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// check
func ValidateToken(tokenString string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
