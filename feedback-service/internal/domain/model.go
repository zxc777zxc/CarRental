package domain

import "time"

type Feedback struct {
	ID        int64
	RentalID  int64
	UserID    int64
	Rating    int32
	Comment   string
	CreatedAt time.Time
}

type FeedbackRepository interface {
	CreateFeedback(f *Feedback) (int64, error)
	GetFeedbackByRental(rentalID int64) ([]*Feedback, error)
}
