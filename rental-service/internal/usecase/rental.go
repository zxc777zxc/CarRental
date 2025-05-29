package usecase

import (
	"CarRental/rental-service/internal/domain"
	"context"
)

type RentalRepo interface {
	Create(ctx context.Context, rental *domain.Rental) (int64, error)
	Complete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*domain.Rental, error)
	ListByUser(ctx context.Context, userID int64) ([]*domain.Rental, error)
}

type RentalUsecase struct {
	repo RentalRepo
}

func NewRentalUsecase(r RentalRepo) *RentalUsecase {
	return &RentalUsecase{repo: r}
}

func (u *RentalUsecase) RentCar(ctx context.Context, rental *domain.Rental) (int64, error) {
	days := rental.EndDate.Sub(rental.StartDate).Hours() / 24
	if days < 1 {
		days = 1
	}
	rental.TotalCost = days * rental.TotalCost // Стоимость за день передаётся заранее
	rental.Status = "active"
	return u.repo.Create(ctx, rental)
}

func (u *RentalUsecase) CompleteRental(ctx context.Context, id int64) error {
	return u.repo.Complete(ctx, id)
}

func (u *RentalUsecase) GetRental(ctx context.Context, id int64) (*domain.Rental, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *RentalUsecase) ListUserRentals(ctx context.Context, userID int64) ([]*domain.Rental, error) {
	return u.repo.ListByUser(ctx, userID)
}
