package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

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
