package usecase

import (
	"CarRental/user-service/internal/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) Create(ctx context.Context, u *domain.User) (int64, error) {
	args := m.Called(ctx, u)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockUserRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	repo := new(mockUserRepo)
	uc := NewUserUsecase(repo)

	user := &domain.User{Email: "test@example.com", Name: "John", Phone: "123456"}
	repo.On("Create", mock.Anything, user).Return(int64(1), nil)

	id, err := uc.Create(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}
