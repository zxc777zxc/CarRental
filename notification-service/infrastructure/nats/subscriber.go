package nats

import (
	"CarRental/notification-service/infrastructure/email"
	"CarRental/notification-service/internal/domain"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func Subscribe(nc *nats.Conn, sender *email.EmailSender) {
	nc.Subscribe("ap2.notification.email", func(m *nats.Msg) {
		var notification domain.EmailNotification
		if err := json.Unmarshal(m.Data, &notification); err != nil {
			log.Println("Invalid email event:", err)
			return
		}
		if err := sender.SendEmail(notification); err != nil {
			log.Println("Email sending error:", err)
		} else {
			log.Println("Email sent to:", notification.To)
		}
	})
}
