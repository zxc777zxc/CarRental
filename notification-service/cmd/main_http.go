package main

import (
	"CarRental/notification-service/config"
	"CarRental/notification-service/infrastructure/email"
	"CarRental/notification-service/internal/domain"
	"encoding/json"
	"log"
	"net/http"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	cfg := config.Load()
	sender := email.NewEmailSender(cfg)

	http.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req EmailRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		emailReq := domain.EmailNotification{
			To:      req.To,
			Subject: req.Subject,
			Body:    req.Body,
		}

		if err := sender.SendEmail(emailReq); err != nil {
			log.Println("Failed to send email:", err)
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email sent successfully"))
	})

	log.Println("HTTP Email API listening on :50058")
	http.ListenAndServe(":50058", nil)
}
