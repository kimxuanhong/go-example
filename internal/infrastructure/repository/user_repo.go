package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/kimxuanhong/go-database/repo"
	"time"

	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
)

const (
	userCacheKeyPrefix = "user:"
	userCacheDuration  = 3 * time.Minute
)

type userRepo struct {
	db    *repo.Repository[UserModel, string]
	repDB *repo.Repository[UserModel, string]
}

func NewUserRepo(db pkg.MainPostgres, repDB pkg.ReplicaPostgres) domain.UserRepository {
	return &userRepo{
		db:    repo.NewRepository[UserModel, string](db),
		repDB: repo.NewRepository[UserModel, string](repDB),
	}
}

func (r *userRepo) GetByUsername(ctx context.Context, userName string) (*domain.User, error) {
	user, err := r.db.FindByID(ctx, "78c83478-5e15-4720-9acb-b70ab32f011b")
	if err != nil {
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
	userModel := &UserModel{
		ID:       uuid.NewString(),
		UserName: user.UserName,
	}

	if err := r.repDB.Insert(ctx, userModel); err != nil {
		return nil, err
	}

	return user, nil
}
