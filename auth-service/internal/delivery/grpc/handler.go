package grpc

import (
	"CarRental/auth-service/internal/usecase"
	pb "CarRental/auth-service/proto"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	token, err := h.usecase.Register(ctx, req.Email, req.Password)
	return &pb.AuthResponse{Token: token}, err
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, err := h.usecase.Login(ctx, req.Email, req.Password)
	return &pb.AuthResponse{Token: token}, err
}

func (h *AuthHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	email, err := h.usecase.Validate(req.Token)
	if err != nil {
		return &pb.ValidateResponse{Valid: false}, nil
	}
	return &pb.ValidateResponse{Valid: true, Email: email}, nil
}
