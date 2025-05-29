package nats

import (
	"CarRental/statistics-service/internal/usecase"
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type StatisticsEvent struct {
	EventType string  `json:"event_type"` // "order" or "feedback"
	Amount    float64 `json:"amount,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
}

func Subscribe(nc *nats.Conn, uc *usecase.StatisticsUsecase) {
	nc.Subscribe("ap2.statistics.event.updated", func(m *nats.Msg) {
		var evt StatisticsEvent
		if err := json.Unmarshal(m.Data, &evt); err != nil {
			log.Println("NATS event parse error:", err)
			return
		}
		switch evt.EventType {
		case "order":
			_ = uc.IncrementOrder(context.Background(), evt.Amount)
		case "feedback":
			_ = uc.AddFeedback(context.Background(), evt.Rating)
		}
	})
}
