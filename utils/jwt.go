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

// NewJWTManager creates a new instance of JWTManager with the provided secret key and token duration.
