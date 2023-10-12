package utils

import (
	"regexp"
)

// Validator defines a struct for input validation.
type Validator struct {
	regexEmail *regexp.Regexp
}

// NewValidator creates a new instance of Validator.
func NewValidator() *Validator {
	regexEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	return &Validator{
		regexEmail: regexEmail,
	}
}

// ValidateEmail checks if the provided email is valid.
func (v *Validator) ValidateEmail(email string) bool {
	return v.regexEmail.MatchString(email)
}
