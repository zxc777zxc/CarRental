package usecase

import (
	"CarRental/feedback-service/internal/domain"
	"context"
	"time"
)

type FeedbackUsecase struct {
	repo domain.FeedbackRepository
}

func NewFeedbackUsecase(r domain.FeedbackRepository) *FeedbackUsecase {
	return &FeedbackUsecase{repo: r}
}

func (u *FeedbackUsecase) SubmitFeedback(ctx context.Context, f *domain.Feedback) (int64, error) {
	f.CreatedAt = time.Now()
	return u.repo.CreateFeedback(f)
}

func (u *FeedbackUsecase) GetByRental(ctx context.Context, rentalID int64) ([]*domain.Feedback, error) {
	return u.repo.GetFeedbackByRental(rentalID)
}
