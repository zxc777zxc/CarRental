package usecase

import (
	"CarRental/auth-service/infrastructure/token"
	"CarRental/auth-service/internal/domain"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) Register(ctx context.Context, email, password string) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{Email: email, Password: string(hashed)}
	err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return token.GenerateToken(email)
}

func (uc *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	return token.GenerateToken(email)
}

func (uc *AuthUsecase) Validate(tokenStr string) (string, error) {
	return token.ValidateToken(tokenStr)
}
