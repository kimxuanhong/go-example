package usecase

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (uc *UserUsecase) GetUser(ctx context.Context, userName string) (*domain.User, error) {
	return uc.repo.GetByUsername(ctx, userName)
}
