package usecase

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/internal/domain/errors"
	"github.com/kimxuanhong/go-example/internal/domain/validator"
	"github.com/kimxuanhong/go-example/internal/infrastructure/external"
)

type UserUsecase struct {
	repo           domain.UserRepository
	accountClient  *external.AccountClient
	consumerClient *external.ConsumerClient
	validator      validator.UserValidator
}

func NewUserUsecase(
	repo domain.UserRepository,
	accountClient *external.AccountClient,
	consumerClient *external.ConsumerClient,
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

	// Kiểm tra user trong account service
	user, err := uc.accountClient.GetUser(ctx, userName, &external.GetUserOptions{
		Fields: []string{"username"},
	})
	if err != nil {
		return nil, errors.NewDomainError("EXTERNAL_ERROR", "failed to check user in account service", err)
	}
	if user == nil {
		return nil, errors.NewDomainError("NOT_FOUND", "user not found in account service", errors.ErrNotFound)
	}

	// Lấy user từ local database
	localUser, err := uc.repo.GetByUsername(ctx, userName)
	if err != nil {
		return nil, errors.NewDomainError("NOT_FOUND", "user not found", errors.ErrNotFound)
	}

	return localUser, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.validateAndCheckDuplicate(ctx, user); err != nil {
		return nil, err
	}

	// Validate trong external systems
	if err := uc.validateInExternalSystems(ctx, user); err != nil {
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

func (uc *UserUsecase) validateInExternalSystems(ctx context.Context, user *domain.User) error {
	// Kiểm tra user trong account service
	existingUser, err := uc.accountClient.GetUser(ctx, user.UserName, &external.GetUserOptions{
		Fields: []string{"username"},
	})
	if err != nil {
		return errors.NewDomainError("EXTERNAL_ERROR", "failed to check user in account service", err)
	}
	if existingUser != nil {
		return errors.NewDomainError("DUPLICATE", "user already exists in account service", errors.ErrValidation)
	}

	// Kiểm tra thông tin consumer
	consumerInfo, err := uc.consumerClient.GetConsumerInfo(ctx, user.UserName, &external.GetConsumerInfoOptions{
		IncludeMetadata: true,
	})
	if err != nil {
		return errors.NewDomainError("EXTERNAL_ERROR", "failed to get consumer info", err)
	}
	if consumerInfo != nil && !consumerInfo.IsActive {
		return errors.NewDomainError("INVALID_STATE", "consumer is not active", errors.ErrValidation)
	}

	return nil
}
