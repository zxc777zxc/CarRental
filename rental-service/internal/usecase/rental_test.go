package usecase

import (
    "CarRental/rental-service/internal/domain"
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type mockRentalRepo struct{ mock.Mock }

func (m *mockRentalRepo) Create(ctx context.Context, r *domain.Rental) (int64, error) {
    args := m.Called(ctx, r)
    return args.Get(0).(int64), args.Error(1)
}

func (m *mockRentalRepo) Complete(ctx context.Context, rentalID int64) error {
    args := m.Called(ctx, rentalID)
    return args.Error(0)
}

func (m *mockRentalRepo) Get(ctx context.Context, id int64) (*domain.Rental, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*domain.Rental), args.Error(1)
}

func TestCreateRental(t *testing.T) {
    repo := new(mockRentalRepo)
    uc := NewRentalUsecase(repo)

    rental := &domain.Rental{UserID: 1, CarID: 2}
    repo.On("Create", mock.Anything, rental).Return(int64(1), nil)

    id, err := uc.Create(context.Background(), rental)
    assert.NoError(t, err)
    assert.Equal(t, int64(1), id)
}
