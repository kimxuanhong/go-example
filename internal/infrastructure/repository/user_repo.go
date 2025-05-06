package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kimxuanhong/go-example/internal/domain"
	"github.com/kimxuanhong/go-example/pkg"
	"github.com/kimxuanhong/go-redis/redis"
)

const (
	userCacheKeyPrefix = "user:"
	userCacheDuration  = 3 * time.Minute
)

type userRepo struct {
	db    pkg.MainPostgres
	repDB pkg.ReplicaPostgres
	redis redis.Redis
}

func NewUserRepo(db pkg.MainPostgres, repDB pkg.ReplicaPostgres, redis redis.Redis) domain.UserRepository {
	return &userRepo{db, repDB, redis}
}

func (r *userRepo) GetByUsername(ctx context.Context, userName string) (*domain.User, error) {
	// Try to get from cache first
	cacheKey := fmt.Sprintf("%s%s", userCacheKeyPrefix, userName)
	cachedUser, err := r.getFromCache(ctx, cacheKey)
	if err == nil && cachedUser != nil {
		return cachedUser, nil
	}

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

	// Cache the result
	if err := r.setCache(ctx, cacheKey, domainUser); err != nil {
		// Log error but don't return it since we already have the data
		fmt.Printf("Failed to cache user data: %v\n", err)
	}

	return domainUser, nil
}

func (r *userRepo) Store(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.repDB.Insert(ctx, user); err != nil {
		return nil, err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("%s%s", userCacheKeyPrefix, user.UserName)
	if err := r.redis.Delete(ctx, cacheKey); err != nil {
		fmt.Printf("Failed to invalidate cache: %v\n", err)
	}

	return user, nil
}

func (r *userRepo) getFromCache(ctx context.Context, key string) (*domain.User, error) {
	data, err := r.redis.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var user domain.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) setCache(ctx context.Context, key string, user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, key, string(data))
}
