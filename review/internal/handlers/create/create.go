package create

import "review/internal/domain/model"

type ReviewCreater interface {
	CreateReview(review *model.Review) (int64, error)
}
