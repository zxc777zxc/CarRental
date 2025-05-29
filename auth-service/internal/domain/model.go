package domain

import (
	"context"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
