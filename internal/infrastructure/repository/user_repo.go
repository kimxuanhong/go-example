package repository

import (
	"context"
	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

type userRepo struct {
	db    pkg.MainPostgres
	repDB pkg.ReplicaPostgres
}

func NewUserRepo(db pkg.MainPostgres, repDB pkg.ReplicaPostgres) domain.UserRepository {
	return &userRepo{db, repDB}
}

func (r *userRepo) GetByUsername(ctx context.Context, userName string) (*domain.User, error) {
	var user UserModel
	if err := r.db.SelectOne(ctx, &user, "user_name = ?", userName); err != nil {
		return nil, err
	}
	return &domain.User{ID: user.ID, UserName: user.UserName}, nil
}

func (r *userRepo) Store(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}
