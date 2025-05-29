package usecase

import (
	"CarRental/user-service/internal/domain"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (int64, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) Create(ctx context.Context, user *domain.User) (int64, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *UserUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserUsecase) Update(ctx context.Context, user *domain.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.DeleteUser(ctx, id)
}
