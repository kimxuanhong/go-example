package rules

import (
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/internal/domain/errors"
)

type UserValidator interface {
	Validate(user *domain.User) error
}

type userValidator struct{}

func NewUserValidator() UserValidator {
	return &userValidator{}
}

func (v *userValidator) Validate(user *domain.User) error {
	if user == nil {
		return errors.NewDomainError("VALIDATION_ERROR", "user cannot be nil", errors.ErrValidation)
	}

	if user.UserName == "" {
		return errors.NewDomainError("VALIDATION_ERROR", "username is required", errors.ErrValidation)
	}

	if user.Email != "" {
		// Add email validation logic here
		// For example, you can use regex or a validation library
	}

	if user.Status != "" && user.Status != "active" && user.Status != "inactive" {
		return errors.NewDomainError("VALIDATION_ERROR", "invalid status value", errors.ErrValidation)
	}

	return nil
}
