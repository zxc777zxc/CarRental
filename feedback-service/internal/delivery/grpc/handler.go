package grpc

import (
	"CarRental/feedback-service/internal/domain"
	"CarRental/feedback-service/internal/usecase"
	pb "CarRental/feedback-service/proto"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedbackHandler struct {
	pb.UnimplementedFeedbackServiceServer
	usecase *usecase.FeedbackUsecase
}

func NewFeedbackHandler(u *usecase.FeedbackUsecase) *FeedbackHandler {
	return &FeedbackHandler{usecase: u}
}

func (h *FeedbackHandler) SubmitFeedback(ctx context.Context, req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error) {
	f := &domain.Feedback{
		RentalID: req.RentalId,
		UserID:   req.UserId,
		Rating:   req.Rating,
		Comment:  req.Comment,
	}
	id, err := h.usecase.SubmitFeedback(ctx, f)
	if err != nil {
		return nil, err
	}
	return &pb.SubmitFeedbackResponse{FeedbackId: id}, nil
}

func (h *FeedbackHandler) GetFeedbackByRental(ctx context.Context, req *pb.GetFeedbackByRentalRequest) (*pb.GetFeedbackByRentalResponse, error) {
	feedbacks, err := h.usecase.GetByRental(ctx, req.RentalId)
	if err != nil {
		return nil, err
	}
	var pbFeedbacks []*pb.Feedback
	for _, f := range feedbacks {
		pbFeedbacks = append(pbFeedbacks, &pb.Feedback{
			Id:        f.ID,
			RentalId:  f.RentalID,
			UserId:    f.UserID,
			Rating:    f.Rating,
			Comment:   f.Comment,
			CreatedAt: timestamppb.New(f.CreatedAt),
		})
	}
	return &pb.GetFeedbackByRentalResponse{Feedbacks: pbFeedbacks}, nil
}
