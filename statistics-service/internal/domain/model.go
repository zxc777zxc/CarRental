package domain

type Statistics struct {
	TotalOrders    int64
	TotalRevenue   float64
	AverageRating  float64
	TotalFeedbacks int64
}

type StatisticsRepository interface {
	IncrementOrder(total float64) error
	AddFeedback(rating float64) error
	GetStatistics() (*Statistics, error)
}
