package usecase

import (
	"CarRental/feedback-service/internal/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockFeedbackRepo struct{ mock.Mock }

func (m *mockFeedbackRepo) Create(ctx context.Context, f *domain.Feedback) (int64, error) {
	args := m.Called(ctx, f)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockFeedbackRepo) GetByRentalID(ctx context.Context, id int64) (*domain.Feedback, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Feedback), args.Error(1)
}

func TestSubmitFeedback(t *testing.T) {
	repo := new(mockFeedbackRepo)
	uc := NewFeedbackUsecase(repo)

	feedback := &domain.Feedback{UserID: 1, RentalID: 1, Rating: 5, Comment: "Great"}
	repo.On("Create", mock.Anything, feedback).Return(int64(1), nil)

	id, err := uc.SubmitFeedback(context.Background(), feedback)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}
