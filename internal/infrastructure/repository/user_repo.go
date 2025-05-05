package repository

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-postgres/postgres"
)

type userRepo struct {
	db postgres.Postgres
}

func NewUserRepo(db postgres.Postgres) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetByUsername(ctx context.Context, userName string) (*domain.User, error) {
	var user UserModel
	if err := r.db.SelectOne(ctx, &user, "user_name = ?", userName); err != nil {
		return nil, err
	}
	return &domain.User{ID: user.ID, UserName: user.UserName}, nil
}
