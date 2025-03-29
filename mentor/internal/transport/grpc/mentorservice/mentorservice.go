package mentorservice

import (
	"context"
	"fmt"
	"mentor/internal/domain/requests"
	client "mentor/pkg/api/proto"
)

const (
	ActionUpdate = "updated"
	ActionDelete = "deleted"
)

type PostgresRepository interface {
	UpdateMentor(ctx context.Context, mentor *requests.RatingRequest) error
	DeleteReviewByMentor(ctx context.Context, mentor *requests.RatingRequest) error
	CreateMentor(ctx context.Context, mentor *requests.MentorRequest) error
}

type MentorService struct {
	client.UnimplementedMentorServiceServer
	repo PostgresRepository
}

func NewMentorService(repo PostgresRepository) *MentorService {
	return &MentorService{repo: repo}
}

func (s *MentorService) MethodMentorRating(ctx context.Context, req *client.RatingRequest) (*client.Response, error) {
	request := &requests.RatingRequest{
		MentorEmail: req.MentorEmail,
		Rating:      req.Rating,
	}

	switch req.Action {
	case ActionDelete:
		err := s.repo.DeleteReviewByMentor(ctx, request)
		if err != nil {
			return &client.Response{
					Success: false,
					Message: "error",
				},
				fmt.Errorf("failed to delete review: %w", err)
		}

		return &client.Response{
			Success: true,
			Message: "ok",
		}, nil

	case ActionUpdate:
		err := s.repo.UpdateMentor(ctx, request)
		if err != nil {
			return &client.Response{
					Success: false,
					Message: "error",
				},
				fmt.Errorf("failed to update review: %w", err)
		}

		return &client.Response{
			Success: true,
			Message: "ok",
		}, nil

	default:
		return &client.Response{
				Success: false,
				Message: "error: action don't matched",
			},
			nil

	}
}

func (s *MentorService) NewMentor(ctx context.Context, req *client.MentorRequest) (*client.Response, error) {
	request := &requests.MentorRequest{
		MentorEmail: req.MentorEmail,
		Contact:     req.Contact,
	}

	err := s.repo.CreateMentor(ctx, request)
	if err != nil {
		return &client.Response{
				Success: false,
				Message: "error",
			},
			fmt.Errorf("failed to create mentor %w", err)
	}

	return &client.Response{
		Success: true,
		Message: "ok",
	}, nil
}
