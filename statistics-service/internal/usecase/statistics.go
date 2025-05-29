package usecase

import (
	"CarRental/statistics-service/internal/domain"
	"context"
)

type StatisticsUsecase struct {
	repo domain.StatisticsRepository
}

func NewStatisticsUsecase(r domain.StatisticsRepository) *StatisticsUsecase {
	return &StatisticsUsecase{repo: r}
}

func (u *StatisticsUsecase) IncrementOrder(ctx context.Context, amount float64) error {
	return u.repo.IncrementOrder(amount)
}

func (u *StatisticsUsecase) AddFeedback(ctx context.Context, rating float64) error {
	return u.repo.AddFeedback(rating)
}

func (u *StatisticsUsecase) GetStats(ctx context.Context) (*domain.Statistics, error) {
	return u.repo.GetStatistics()
}
