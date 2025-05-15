package repository

import (
	"context"
	"time"

	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

const (
	userCacheKeyPrefix = "user:"
	userCacheDuration  = 3 * time.Minute
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

	domainUser := &domain.User{
		ID:        user.ID,
		PartnerId: user.PartnerId,
		Total:     user.Total,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Status:    user.Status,
	}

	return domainUser, nil
}

func (r *userRepo) Store(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.repDB.Insert(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
