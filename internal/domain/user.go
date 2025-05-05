package domain

import "context"

type User struct {
	ID        string
	PartnerId string
	Total     int
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Status    string
}

type UserRepository interface {
	GetByUsername(ctx context.Context, userName string) (*User, error)
	Store(ctx context.Context, user *User) (*User, error)
}
