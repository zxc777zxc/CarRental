package repository

import (
	"CarRental/auth-service/internal/domain"
	"context"
	"database/sql"
	//"errors"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
