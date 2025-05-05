package usecase

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

type UserUsecase struct {
	repo           domain.UserRepository
	accountClient  pkg.AccountClient
	consumerClient pkg.ConsumerClient
}

func NewUserUsecase(repo domain.UserRepository, accountClient pkg.AccountClient, consumerClient pkg.ConsumerClient) *UserUsecase {
	return &UserUsecase{repo, accountClient, consumerClient}
}

func (uc *UserUsecase) GetUser(ctx context.Context, userName string) (*domain.User, error) {
	return uc.repo.GetByUsername(ctx, userName)
}
