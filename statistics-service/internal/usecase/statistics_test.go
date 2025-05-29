package usecase

import (
    "CarRental/statistics-service/internal/domain"
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
)

type mockStatsRepo struct{}

func (m *mockStatsRepo) IncrementOrderCount(ctx context.Context) error { return nil }
func (m *mockStatsRepo) AddFeedback(ctx context.Context, rating int32) error {
    if rating < 1 || rating > 5 {
        return domain.ErrInvalidRating
    }
    return nil
}

func TestAddFeedback(t *testing.T) {
    repo := &mockStatsRepo{}
    uc := NewStatisticsUsecase(repo)

    err := uc.AddFeedback(context.Background(), 5)
    assert.NoError(t, err)

    err = uc.AddFeedback(context.Background(), 0)
    assert.Error(t, err)
}
