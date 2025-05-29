package grpc

import (
	"CarRental/notification-service/infrastructure/email"
	"CarRental/notification-service/internal/domain"
	pb "CarRental/notification-service/proto"
	"context"
	"log"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	sender *email.EmailSender
}

func NewNotificationHandler(sender *email.EmailSender) *NotificationHandler {
	return &NotificationHandler{sender: sender}
}

func (h *NotificationHandler) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	email := domain.EmailNotification{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}

	err := h.sender.SendEmail(email)
	if err != nil {
		log.Printf("Email sending failed: %v", err)
		return nil, err
	}

	return &pb.EmailResponse{Message: "Email sent successfully"}, nil
}
