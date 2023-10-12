package utils

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
