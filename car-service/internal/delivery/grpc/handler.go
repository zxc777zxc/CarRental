package grpc

import (
	"CarRental/car-service/internal/domain"
	"CarRental/car-service/internal/usecase"
	pb "CarRental/car-service/proto"
	"context"
)

type CarHandler struct {
	pb.UnimplementedCarServiceServer
	uc *usecase.CarUsecase
}

func NewCarHandler(uc *usecase.CarUsecase) *CarHandler {
	return &CarHandler{uc: uc}
}

func (h *CarHandler) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CarResponse, error) {
	car := &domain.Car{
		Brand:        req.Brand,
		Model:        req.Model,
		Fuel:         req.Fuel,
		Transmission: req.Transmission,
		PricePerDay:  req.PricePerDay,
	}
	id, err := h.uc.Create(ctx, car)
	if err != nil {
		return nil, err
	}
	car.ID = id
	return &pb.CarResponse{Car: toProtoCar(car)}, nil
}

func (h *CarHandler) GetCarByID(ctx context.Context, req *pb.GetCarByIDRequest) (*pb.CarResponse, error) {
	car, err := h.uc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CarResponse{Car: toProtoCar(car)}, nil
}

func (h *CarHandler) ListCars(ctx context.Context, _ *pb.Empty) (*pb.CarListResponse, error) {
	cars, err := h.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	var protoCars []*pb.Car
	for _, car := range cars {
		protoCars = append(protoCars, toProtoCar(car))
	}
	return &pb.CarListResponse{Cars: protoCars}, nil
}

func (h *CarHandler) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.CarResponse, error) {
	car := &domain.Car{
		ID:           req.Id,
		Brand:        req.Brand,
		Model:        req.Model,
		Fuel:         req.Fuel,
		Transmission: req.Transmission,
		PricePerDay:  req.PricePerDay,
	}
	err := h.uc.Update(ctx, car)
	if err != nil {
		return nil, err
	}
	return &pb.CarResponse{Car: toProtoCar(car)}, nil
}

func (h *CarHandler) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteResponse, error) {
	err := h.uc.Delete(ctx, req.Id)
	return &pb.DeleteResponse{Success: err == nil}, err
}

func toProtoCar(c *domain.Car) *pb.Car {
	return &pb.Car{
		Id:           c.ID,
		Brand:        c.Brand,
		Model:        c.Model,
		Fuel:         c.Fuel,
		Transmission: c.Transmission,
		PricePerDay:  c.PricePerDay,
	}
}
