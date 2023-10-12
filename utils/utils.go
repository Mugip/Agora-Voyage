package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AppError is a custom error type that provides additional information about application errors.
type AppError struct {
	Message string
	Code    int
}

// Error returns the error message as a string.
func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new instance of AppError with the provided error message and code.
func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

// IsAppError checks if the given error is an instance of AppError.
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// WrapError wraps the given error with additional context information.
func WrapError(err error, message string, code int) *AppError {
	return &AppError{
		Message: message + ": " + err.Error(),
		Code:    code,
	}
}

// ErrNotFound is an app-specific error instance indicating a "not found" error.
var ErrNotFound = NewAppError("Not Found", 404)

// ErrUnauthorized is an app-specific error instance indicating an "unauthorized" error.
var ErrUnauthorized = NewAppError("Unauthorized", 401)

// ErrBadRequest is an app-specific error instance indicating a "bad request" error.
var ErrBadRequest = NewAppError("Bad Request", 400)

// JWTManager defines a struct for JWT management.
type JWTManager struct {
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

// GenerateToken generates a new JWT token with the provided claims.
func (jm *JWTManager) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", WrapError(err, "failed to sign JWT token", 500)
	}
	return signedToken, nil
}

// ParseToken parses and verifies the JWT token and returns the claims.
func (jm *JWTManager) ParseToken(tokenString string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jm.secretKey, nil
	})
	if err != nil {
		return WrapError(err, "failed to parse JWT token", 401)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

// HashPassword generates a bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", WrapError(err, "failed to hash password", 500)
	}
	return string(hashedPassword), nil
}

// ComparePasswordHash checks if the provided password matches the hash.
func ComparePasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return WrapError(err, "invalid password", 401)
	}
	return nil
}
