package usecase

import (
	"context"

	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/internal/domain/errors"
	"github.com/kimxuanhong/go-example/internal/domain/validator"
	"github.com/kimxuanhong/go-example/pkg"
)

type UserUsecase struct {
	repo           domain.UserRepository
	accountClient  pkg.AccountClient
	consumerClient pkg.ConsumerClient
	validator      validator.UserValidator
}

func NewUserUsecase(
	repo domain.UserRepository,
	accountClient pkg.AccountClient,
	consumerClient pkg.ConsumerClient,
	validator validator.UserValidator,
) *UserUsecase {
	return &UserUsecase{
		repo:           repo,
		accountClient:  accountClient,
		consumerClient: consumerClient,
		validator:      validator,
	}
}

func (uc *UserUsecase) GetUser(ctx context.Context, userName string) (*domain.User, error) {
	if userName == "" {
		return nil, errors.NewDomainError("VALIDATION_ERROR", "username is required", errors.ErrValidation)
	}

	user, err := uc.repo.GetByUsername(ctx, userName)
	if err != nil {
		return nil, errors.NewDomainError("NOT_FOUND", "user not found", errors.ErrNotFound)
	}

	return user, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.validateAndCheckDuplicate(ctx, user); err != nil {
		return nil, err
	}

	createdUser, err := uc.repo.Store(ctx, user)
	if err != nil {
		return nil, errors.NewDomainError("INTERNAL_ERROR", "failed to create user", errors.ErrInternal)
	}

	return createdUser, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.validateAndCheckExists(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := uc.repo.Store(ctx, user)
	if err != nil {
		return nil, errors.NewDomainError("INTERNAL_ERROR", "failed to update user", errors.ErrInternal)
	}

	return updatedUser, nil
}

// Helper functions
func (uc *UserUsecase) validateAndCheckDuplicate(ctx context.Context, user *domain.User) error {
	if err := uc.validator.Validate(user); err != nil {
		return err
	}

	existingUser, err := uc.repo.GetByUsername(ctx, user.UserName)
	if err == nil && existingUser != nil {
		return errors.NewDomainError("DUPLICATE", "user already exists", errors.ErrValidation)
	}

	return nil
}

func (uc *UserUsecase) validateAndCheckExists(ctx context.Context, user *domain.User) error {
	if err := uc.validator.Validate(user); err != nil {
		return err
	}

	existingUser, err := uc.repo.GetByUsername(ctx, user.UserName)
	if err != nil || existingUser == nil {
		return errors.NewDomainError("NOT_FOUND", "user not found", errors.ErrNotFound)
	}

	return nil
}
