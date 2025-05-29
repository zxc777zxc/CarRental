package domain

import "time"

type Payment struct {
	ID       int64
	RentalID int64
	Amount   float64
	Method   string
	Status   string
	PaidAt   time.Time
}

type PaymentRepository interface {
	CreatePayment(p *Payment) (int64, error)
	GetPayment(id int64) (*Payment, error)
}
