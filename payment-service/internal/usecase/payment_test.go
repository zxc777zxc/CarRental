package usecase

import (
    "CarRental/payment-service/internal/domain"
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type mockPaymentRepo struct{ mock.Mock }

func (m *mockPaymentRepo) Create(ctx context.Context, p *domain.Payment) (int64, error) {
    args := m.Called(ctx, p)
    return args.Get(0).(int64), args.Error(1)
}

func (m *mockPaymentRepo) Get(ctx context.Context, id int64) (*domain.Payment, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*domain.Payment), args.Error(1)
}

func TestCreatePayment(t *testing.T) {
    repo := new(mockPaymentRepo)
    uc := NewPaymentUsecase(repo)

    payment := &domain.Payment{RentalID: 1, Amount: 100}
    repo.On("Create", mock.Anything, payment).Return(int64(1), nil)

    id, err := uc.Create(context.Background(), payment)
    assert.NoError(t, err)
    assert.Equal(t, int64(1), id)
}
