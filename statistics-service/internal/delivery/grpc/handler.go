package grpc

import (
	"CarRental/statistics-service/internal/usecase"
	pb "CarRental/statistics-service/proto"
	"context"
)

type Handler struct {
	pb.UnimplementedStatisticsServiceServer
	uc *usecase.StatisticsUsecase
}

func NewHandler(u *usecase.StatisticsUsecase) *Handler {
	return &Handler{uc: u}
}

func (h *Handler) GetStatistics(ctx context.Context, req *pb.GetStatisticsRequest) (*pb.GetStatisticsResponse, error) {
	stats, err := h.uc.GetStats(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetStatisticsResponse{
		Stats: &pb.Statistics{
			TotalOrders:    stats.TotalOrders,
			TotalRevenue:   stats.TotalRevenue,
			AverageRating:  stats.AverageRating,
			TotalFeedbacks: stats.TotalFeedbacks,
		},
	}, nil
}
