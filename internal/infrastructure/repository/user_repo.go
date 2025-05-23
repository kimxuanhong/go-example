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

type UserRepository struct {
	*repo.Repository[UserModel, string]
	FindByUserName func(ctx context.Context, username string) (*UserModel, error) `repo:"@Query"`
}

type userRepo struct {
	db    *UserRepository
	repDB *UserRepository
}

func NewUserRepo(db pkg.MainPostgres, repDB pkg.ReplicaPostgres) domain.UserRepository {
	repositoryMain := repo.NewRepository[UserModel, string](db)
	userRepoMain := &UserRepository{Repository: repositoryMain}
	err := repositoryMain.FillFuncFields(userRepoMain)
	if err != nil {
		panic(err)
	}

	return &userRepo{
		db: userRepoMain,
		repDB: &UserRepository{
			Repository: repo.NewRepository[UserModel, string](repDB),
		},
	}
}

func (r *userRepo) GetByUsername(ctx context.Context, userName string) (*domain.User, error) {
	user, err := r.db.FindByUserName(ctx, userName)
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
