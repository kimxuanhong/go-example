package dto

import "github.com/kimxuanhong/go-example/internal/domain"

type UserResponse struct {
	ID        string `json:"id"`
	PartnerId string `json:"partner_id,omitempty"`
	Total     int    `json:"total,omitempty"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Status    string `json:"status,omitempty"`
}

type UserRequest struct {
	UserName  string `json:"user_name" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"omitempty,email"`
	Status    string `json:"status"`
}

func ToUserResponse(user *domain.User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:        user.ID,
		PartnerId: user.PartnerId,
		Total:     user.Total,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Status:    user.Status,
	}
}

func ToUserDomain(req *UserRequest) *domain.User {
	if req == nil {
		return nil
	}
	return &domain.User{
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Status:    req.Status,
	}
}
