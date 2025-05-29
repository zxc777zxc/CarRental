package repository

import (
	"CarRental/user-service/internal/domain"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, u *domain.User) (int64, error) {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO users (email, name, phone) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Name, u.Phone,
	).Scan(&u.ID)
	return u.ID, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, email, name, phone FROM users WHERE id = $1", id,
	)
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Phone)
	return u, err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, email, name, phone FROM users WHERE email = $1", email,
	)
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Phone)
	return u, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, u *domain.User) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET name=$1, phone=$2 WHERE id=$3",
		u.Name, u.Phone, u.ID,
	)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}
