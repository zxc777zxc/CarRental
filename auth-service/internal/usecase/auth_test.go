package usecase

import (
	"CarRental/auth-service/internal/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type mockAuthRepo struct{ mock.Mock }

func (m *mockAuthRepo) CreateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockAuthRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	repo := new(mockAuthRepo)
	uc := NewAuthUsecase(repo)

	email := "user@example.com"
	password := "123456"
	repo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	token, err := uc.Register(context.Background(), email, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLogin_Success(t *testing.T) {
	repo := new(mockAuthRepo)
	uc := NewAuthUsecase(repo)

	email := "user@example.com"
	password := "123456"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{Email: email, Password: string(hashed)}

	repo.On("GetUserByEmail", mock.Anything, email).Return(user, nil)

	token, err := uc.Login(context.Background(), email, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
