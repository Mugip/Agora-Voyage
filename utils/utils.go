package utils

// ErrNotFound is an app-specific error instance indicating a "not found" error.
var ErrNotFound = NewAppError("Not Found", 404)

// ErrUnauthorized is an app-specific error instance indicating an "unauthorized" error.
var ErrUnauthorized = NewAppError("Unauthorized", 401)

// ErrBadRequest is an app-specific error instance indicating a "bad request" error.
var ErrBadRequest = NewAppError("Bad Request", 400)
