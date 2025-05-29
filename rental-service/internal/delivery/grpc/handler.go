package grpc

import (
	"context"

	"CarRental/rental-service/internal/domain"
	"CarRental/rental-service/internal/usecase"
	pb "CarRental/rental-service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RentalHandler struct {
	pb.UnimplementedRentalServiceServer
	usecase *usecase.RentalUsecase
}

func NewRentalHandler(u *usecase.RentalUsecase) *RentalHandler {
	return &RentalHandler{usecase: u}
}

func (h *RentalHandler) RentCar(ctx context.Context, req *pb.RentCarRequest) (*pb.RentCarResponse, error) {
	rental := &domain.Rental{
		UserID:    req.UserId,
		CarID:     req.CarId,
		StartDate: req.StartDate.AsTime(),
		EndDate:   req.EndDate.AsTime(),
		TotalCost: req.DailyPrice,
	}
	id, err := h.usecase.RentCar(ctx, rental)
	if err != nil {
		return nil, err
	}
	return &pb.RentCarResponse{RentalId: id}, nil
}

func (h *RentalHandler) CompleteRental(ctx context.Context, req *pb.CompleteRentalRequest) (*pb.CompleteRentalResponse, error) {
	err := h.usecase.CompleteRental(ctx, req.RentalId)
	if err != nil {
		return nil, err
	}
	return &pb.CompleteRentalResponse{Message: "Rental completed successfully"}, nil
}

func (h *RentalHandler) GetRental(ctx context.Context, req *pb.GetRentalRequest) (*pb.GetRentalResponse, error) {
	r, err := h.usecase.GetRental(ctx, req.RentalId)
	if err != nil {
		return nil, err
	}
	return &pb.GetRentalResponse{Rental: toProto(r)}, nil
}

func (h *RentalHandler) ListUserRentals(ctx context.Context, req *pb.ListUserRentalsRequest) (*pb.ListUserRentalsResponse, error) {
	rentals, err := h.usecase.ListUserRentals(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var result []*pb.Rental
	for _, r := range rentals {
		result = append(result, toProto(r))
	}
	return &pb.ListUserRentalsResponse{Rentals: result}, nil
}

func toProto(r *domain.Rental) *pb.Rental {
	return &pb.Rental{
		Id:        r.ID,
		UserId:    r.UserID,
		CarId:     r.CarID,
		StartDate: timestamppb.New(r.StartDate),
		EndDate:   timestamppb.New(r.EndDate),
		TotalCost: r.TotalCost,
		Status:    r.Status,
	}
}
