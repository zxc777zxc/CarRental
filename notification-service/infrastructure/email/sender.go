package email

import (
	"CarRental/notification-service/config"
	"CarRental/notification-service/internal/domain"
	"fmt"
	"net/smtp"
)

type EmailSender struct {
	cfg *config.Config
}

func NewEmailSender(cfg *config.Config) *EmailSender {
	return &EmailSender{cfg: cfg}
}

func (s *EmailSender) SendEmail(n domain.EmailNotification) error {
	auth := smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPassword, s.cfg.SMTPServer)
	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPServer, s.cfg.SMTPPort)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", n.To, n.Subject, n.Body))

	return smtp.SendMail(addr, auth, s.cfg.SMTPUser, []string{n.To}, msg)
}
