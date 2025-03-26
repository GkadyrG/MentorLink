package repository

import (
	"context"
	"rating/internal/domain/models"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveReview(ctx context.Context, review *models.ReviewEvent) error {
	return nil
}
