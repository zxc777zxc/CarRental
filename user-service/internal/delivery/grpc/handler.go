package grpc

import (
	"CarRental/user-service/internal/domain"
	"CarRental/user-service/internal/usecase"
	pb "CarRental/user-service/proto"
	"context"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &domain.User{
		Email: req.Email,
		Name:  req.Name,
		Phone: req.Phone,
	}
	id, err := h.usecase.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return toProto(user), nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.UserResponse, error) {
	user, err := h.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return toProto(user), nil
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	user, err := h.usecase.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return toProto(user), nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := &domain.User{
		ID:    req.Id,
		Name:  req.Name,
		Phone: req.Phone,
	}
	err := h.usecase.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	updated, _ := h.usecase.GetByID(ctx, user.ID)
	return toProto(updated), nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteResponse, error) {
	err := h.usecase.Delete(ctx, req.Id)
	return &pb.DeleteResponse{Success: err == nil}, err
}

func toProto(u *domain.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:    u.ID,
		Email: u.Email,
		Name:  u.Name,
		Phone: u.Phone,
	}
}
