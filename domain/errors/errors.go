package errors

import (
	"errors"
	"fmt"
)

// Domain error types
var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrInternal        = errors.New("internal server error")
	ErrExternalService = errors.New("external service error")
	ErrValidation      = errors.New("validation error")
)

// DomainError represents a domain-specific error
type DomainError struct {
	Code    string
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

// New creates a new domain error
func New(code, message string, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NewValidationError creates a validation error
func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:    "VALIDATION_ERROR",
		Message: message,
		Err:     ErrValidation,
	}
}

// NewExternalServiceError creates an external service error
func NewExternalServiceError(service string, err error) *DomainError {
	return &DomainError{
		Code:    "EXTERNAL_SERVICE_ERROR",
		Message: fmt.Sprintf("error calling %s", service),
		Err:     err,
	}
}
