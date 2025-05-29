package repository

import (
	"CarRental/feedback-service/internal/domain"
	"database/sql"
)

type FeedbackRepo struct {
	db *sql.DB
}

func NewFeedbackRepo(db *sql.DB) *FeedbackRepo {
	return &FeedbackRepo{db: db}
}

func (r *FeedbackRepo) CreateFeedback(f *domain.Feedback) (int64, error) {
	query := `INSERT INTO feedbacks (rental_id, user_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, f.RentalID, f.UserID, f.Rating, f.Comment, f.CreatedAt).Scan(&f.ID)
	return f.ID, err
}

func (r *FeedbackRepo) GetFeedbackByRental(rentalID int64) ([]*domain.Feedback, error) {
	rows, err := r.db.Query(`SELECT id, rental_id, user_id, rating, comment, created_at FROM feedbacks WHERE rental_id = $1`, rentalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*domain.Feedback
	for rows.Next() {
		var f domain.Feedback
		if err := rows.Scan(&f.ID, &f.RentalID, &f.UserID, &f.Rating, &f.Comment, &f.CreatedAt); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, &f)
	}
	return feedbacks, nil
}
