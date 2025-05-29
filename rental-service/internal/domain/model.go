package domain

import "time"

type Rental struct {
	ID        int64
	UserID    int64
	CarID     int64
	StartDate time.Time
	EndDate   time.Time
	TotalCost float64
	Status    string // "active", "completed", etc.
}
