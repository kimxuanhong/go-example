package facade

import (
	"context"

	"github.com/kimxuanhong/go-example/internal/interface/dto"
	"github.com/kimxuanhong/go-example/internal/usecase"
)

// UserFacade provides a simplified interface for user-related operations
type UserFacade struct {
	userUsecase *usecase.UserUsecase
}

// NewUserFacade creates a new instance of UserFacade
func NewUserFacade(userUsecase *usecase.UserUsecase) *UserFacade {
	return &UserFacade{
		userUsecase: userUsecase,
	}
}

// GetUser retrieves a user by username
func (f *UserFacade) GetUser(ctx context.Context, userName string) (*dto.UserResponse, error) {
	user, err := f.userUsecase.GetUser(ctx, userName)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(user), nil
}

// CreateUser creates a new user
func (f *UserFacade) CreateUser(ctx context.Context, req *dto.UserRequest) (*dto.UserResponse, error) {
	user := dto.ToUserDomain(req)
	createdUser, err := f.userUsecase.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(createdUser), nil
}

// UpdateUser updates an existing user
func (f *UserFacade) UpdateUser(ctx context.Context, userName string, req *dto.UserRequest) (*dto.UserResponse, error) {
	user := dto.ToUserDomain(req)
	user.UserName = userName // Ensure username from path is used
	updatedUser, err := f.userUsecase.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(updatedUser), nil
}
