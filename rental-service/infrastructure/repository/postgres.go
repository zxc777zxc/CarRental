package repository

import (
	"CarRental/rental-service/internal/domain"
	"context"
	"database/sql"
)

type RentalRepo struct {
	db *sql.DB
}

func NewRentalRepo(db *sql.DB) *RentalRepo {
	return &RentalRepo{db: db}
}

func (r *RentalRepo) Create(ctx context.Context, rental *domain.Rental) (int64, error) {
	query := `INSERT INTO rentals (user_id, car_id, start_date, end_date, total_cost, status)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRowContext(ctx, query,
		rental.UserID, rental.CarID, rental.StartDate, rental.EndDate, rental.TotalCost, rental.Status,
	).Scan(&rental.ID)
	return rental.ID, err
}

func (r *RentalRepo) Complete(ctx context.Context, id int64) error {
	query := `UPDATE rentals SET status = 'completed' WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *RentalRepo) GetByID(ctx context.Context, id int64) (*domain.Rental, error) {
	var rental domain.Rental
	query := `SELECT id, user_id, car_id, start_date, end_date, total_cost, status FROM rentals WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query).Scan(
		&rental.ID, &rental.UserID, &rental.CarID, &rental.StartDate, &rental.EndDate, &rental.TotalCost, &rental.Status,
	)
	return &rental, err
}

func (r *RentalRepo) ListByUser(ctx context.Context, userID int64) ([]*domain.Rental, error) {
	query := `SELECT id, user_id, car_id, start_date, end_date, total_cost, status FROM rentals WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentals []*domain.Rental
	for rows.Next() {
		var r domain.Rental
		if err := rows.Scan(&r.ID, &r.UserID, &r.CarID, &r.StartDate, &r.EndDate, &r.TotalCost, &r.Status); err != nil {
			return nil, err
		}
		rentals = append(rentals, &r)
	}
	return rentals, nil
}
