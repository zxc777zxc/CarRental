package repository

import (
	"CarRental/car-service/internal/domain"
	"context"
	"database/sql"
)

type CarRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (r *CarRepository) Create(ctx context.Context, car *domain.Car) (int64, error) {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO cars (brand, model, fuel, transmission, price_per_day)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		car.Brand, car.Model, car.Fuel, car.Transmission, car.PricePerDay,
	).Scan(&car.ID)
	return car.ID, err
}

func (r *CarRepository) GetByID(ctx context.Context, id int64) (*domain.Car, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, brand, model, fuel, transmission, price_per_day FROM cars WHERE id=$1`, id,
	)
	car := &domain.Car{}
	err := row.Scan(&car.ID, &car.Brand, &car.Model, &car.Fuel, &car.Transmission, &car.PricePerDay)
	return car, err
}

func (r *CarRepository) List(ctx context.Context) ([]*domain.Car, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, brand, model, fuel, transmission, price_per_day FROM cars`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []*domain.Car
	for rows.Next() {
		car := &domain.Car{}
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Fuel, &car.Transmission, &car.PricePerDay)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *CarRepository) Update(ctx context.Context, car *domain.Car) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE cars SET brand=$1, model=$2, fuel=$3, transmission=$4, price_per_day=$5 WHERE id=$6`,
		car.Brand, car.Model, car.Fuel, car.Transmission, car.PricePerDay, car.ID,
	)
	return err
}

func (r *CarRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM cars WHERE id=$1`, id)
	return err
}
