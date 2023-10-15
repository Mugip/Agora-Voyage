package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTManager defines a struct for JWT management.
type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

// JWTManagerV2 defines a struct for JWT management (version 2).
type JWTManagerV2 struct {
	secretKey     []byte
	tokenDuration time.Duration
}

// NewJWTManager creates a new instance of JWTManager with the provided secret key and token duration.
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(secretKey),
		tokenDuration: tokenDuration,
	}
}

// NewJWTManagerV2 creates a new instance of JWTManagerV2 with the provided secret key and token duration.
func NewJWTManagerV2(secretKey string, tokenDuration time.Duration) *JWTManagerV2 {
	return &JWTManagerV2{
		secretKey:     []byte(secretKey),
		tokenDuration: tokenDuration,
	}
}

// GenerateToken generates a new JWT token with the specified claims.
func (jm *JWTManager) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseToken parses the provided JWT token and returns the claims.
func (jm *JWTManager) ParseToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jm.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// GenerateTokenV2 generates a new JWT token with the specified claims (version 2).
func (jm *JWTManagerV2) GenerateTokenV2(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseTokenV2 parses the provided JWT token and returns the claims (version 2).
func (jm *JWTManagerV2) ParseTokenV2(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jm.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
