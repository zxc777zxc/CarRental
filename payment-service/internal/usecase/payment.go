package usecase

import (
	"CarRental/payment-service/internal/domain"
	"context"
	"time"
)

type PaymentUsecase struct {
	repo domain.PaymentRepository
}

func NewPaymentUsecase(r domain.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{repo: r}
}

func (u *PaymentUsecase) ProcessPayment(ctx context.Context, p *domain.Payment) (int64, error) {
	p.Status = "paid"
	p.PaidAt = time.Now()
	return u.repo.CreatePayment(p)
}

func (u *PaymentUsecase) GetPayment(ctx context.Context, id int64) (*domain.Payment, error) {
	return u.repo.GetPayment(id)
}
