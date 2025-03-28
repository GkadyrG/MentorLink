package mentorservice

import (
	"context"
	"mentor/internal/domain/requests"
	client "mentor/pkg/api/proto"
)

type Repository interface {
	UpdateMentor(ctx context.Context, mentor *requests.RatingRequest) error
	DeleteReviewByMentor(ctx context.Context, mentor *requests.RatingRequest) error
}

type MentorService struct {
	client.UnimplementedMentorServiceServer
	repo Repository
}

func NewOrderService(repo Repository) *MentorService {
	return &MentorService{repo: repo}
}

func (s *MentorService) MethodMentorRating(ctx context.Context, req)