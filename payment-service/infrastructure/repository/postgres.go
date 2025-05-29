package repository

import (
	"CarRental/payment-service/internal/domain"
	"database/sql"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) CreatePayment(p *domain.Payment) (int64, error) {
	query := `INSERT INTO payments (rental_id, amount, method, status, paid_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, p.RentalID, p.Amount, p.Method, p.Status, p.PaidAt).Scan(&p.ID)
	return p.ID, err
}

func (r *PaymentRepo) GetPayment(id int64) (*domain.Payment, error) {
	row := r.db.QueryRow(`SELECT id, rental_id, amount, method, status, paid_at FROM payments WHERE id = $1`, id)
	var p domain.Payment
	err := row.Scan(&p.ID, &p.RentalID, &p.Amount, &p.Method, &p.Status, &p.PaidAt)
	return &p, err
}
