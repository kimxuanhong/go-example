package errors

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrValidation   = errors.New("validation error")
	ErrInternal     = errors.New("internal server error")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type DomainError struct {
	Code    string
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func NewDomainError(code string, message string, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
