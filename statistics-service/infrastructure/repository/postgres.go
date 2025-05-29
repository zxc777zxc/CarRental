package repository

import (
	"CarRental/statistics-service/internal/domain"
	"database/sql"
)

type StatisticsRepo struct {
	db *sql.DB
}

func NewStatisticsRepo(db *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{db: db}
}

func (r *StatisticsRepo) IncrementOrder(amount float64) error {
	_, err := r.db.Exec(`UPDATE statistics SET total_orders = total_orders + 1, total_revenue = total_revenue + $1`, amount)
	return err
}

func (r *StatisticsRepo) AddFeedback(rating float64) error {
	_, err := r.db.Exec(`UPDATE statistics SET total_feedbacks = total_feedbacks + 1, average_rating = 
		((average_rating * (total_feedbacks) + $1) / (total_feedbacks + 1))`, rating)
	return err
}

func (r *StatisticsRepo) GetStatistics() (*domain.Statistics, error) {
	row := r.db.QueryRow(`SELECT total_orders, total_revenue, average_rating, total_feedbacks FROM statistics`)
	var s domain.Statistics
	err := row.Scan(&s.TotalOrders, &s.TotalRevenue, &s.AverageRating, &s.TotalFeedbacks)
	return &s, err
}
