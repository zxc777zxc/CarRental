package grpc

import (
	"context"
	_ "time"

	"CarRental/payment-service/internal/domain"
	"CarRental/payment-service/internal/usecase"
	pb "CarRental/payment-service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	usecase *usecase.PaymentUsecase
}

func NewPaymentHandler(u *usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{usecase: u}
}

func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	p := &domain.Payment{
		RentalID: req.RentalId,
		Amount:   req.Amount,
		Method:   req.Method,
	}
	id, err := h.usecase.ProcessPayment(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.ProcessPaymentResponse{PaymentId: id}, nil
}

func (h *PaymentHandler) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	p, err := h.usecase.GetPayment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetPaymentResponse{Payment: &pb.Payment{
		Id:       p.ID,
		RentalId: p.RentalID,
		Amount:   p.Amount,
		Method:   p.Method,
		Status:   p.Status,
		PaidAt:   timestamppb.New(p.PaidAt),
	}}, nil
}
