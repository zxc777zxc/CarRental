package usecase

import (
	"CarRental/car-service/internal/domain"
	"context"
)

type CarRepository interface {
	Create(context.Context, *domain.Car) (int64, error)
	GetByID(context.Context, int64) (*domain.Car, error)
	List(context.Context) ([]*domain.Car, error)
	Update(context.Context, *domain.Car) error
	Delete(context.Context, int64) error
}

type CarCache interface {
	SetCar(context.Context, *domain.Car) error
	GetCar(context.Context, int64) (*domain.Car, error)
	SetCarList(context.Context, []*domain.Car) error
	GetCarList(context.Context) ([]*domain.Car, error)
}

type CarUsecase struct {
	repo  CarRepository
	cache CarCache
}

func NewCarUsecase(r CarRepository, c CarCache) *CarUsecase {
	return &CarUsecase{repo: r, cache: c}
}

func (u *CarUsecase) Create(ctx context.Context, car *domain.Car) (int64, error) {
	id, err := u.repo.Create(ctx, car)
	if err == nil {
		_ = u.cache.SetCar(ctx, car)
	}
	return id, err
}

func (u *CarUsecase) GetByID(ctx context.Context, id int64) (*domain.Car, error) {
	car, err := u.cache.GetCar(ctx, id)
	if err == nil {
		return car, nil
	}
	car, err = u.repo.GetByID(ctx, id)
	if err == nil {
		_ = u.cache.SetCar(ctx, car)
	}
	return car, err
}

func (u *CarUsecase) List(ctx context.Context) ([]*domain.Car, error) {
	cars, err := u.cache.GetCarList(ctx)
	if err == nil {
		return cars, nil
	}
	cars, err = u.repo.List(ctx)
	if err == nil {
		_ = u.cache.SetCarList(ctx, cars)
	}
	return cars, err
}

func (u *CarUsecase) Update(ctx context.Context, car *domain.Car) error {
	err := u.repo.Update(ctx, car)
	if err == nil {
		_ = u.cache.SetCar(ctx, car)
	}
	return err
}

func (u *CarUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
