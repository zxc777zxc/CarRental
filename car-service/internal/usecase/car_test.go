package usecase

import (
	"CarRental/car-service/internal/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCarRepo struct{ mock.Mock }
type mockCarCache struct{ mock.Mock }

func (m *mockCarRepo) Create(ctx context.Context, car *domain.Car) (int64, error) {
	args := m.Called(ctx, car)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockCarRepo) GetByID(ctx context.Context, id int64) (*domain.Car, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *mockCarRepo) List(ctx context.Context) ([]*domain.Car, error) {
	return nil, nil
}
func (m *mockCarRepo) Update(ctx context.Context, car *domain.Car) error { return nil }
func (m *mockCarRepo) Delete(ctx context.Context, id int64) error        { return nil }

func (m *mockCarCache) SetCar(ctx context.Context, car *domain.Car) error         { return nil }
func (m *mockCarCache) GetCar(ctx context.Context, id int64) (*domain.Car, error) { return nil, nil }
func (m *mockCarCache) SetCarList(ctx context.Context, cars []*domain.Car) error  { return nil }
func (m *mockCarCache) GetCarList(ctx context.Context) ([]*domain.Car, error)     { return nil, nil }

func TestCreateCar(t *testing.T) {
	repo := new(mockCarRepo)
	cache := new(mockCarCache)
	uc := NewCarUsecase(repo, cache)

	car := &domain.Car{Brand: "Toyota", Model: "Corolla", Fuel: "Gas", Transmission: "Auto", PricePerDay: 50.0}
	repo.On("Create", mock.Anything, car).Return(int64(1), nil)

	id, err := uc.Create(context.Background(), car)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}
