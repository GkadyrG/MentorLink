package delete

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

type DelReview interface {
	DeleteReview(id int64) error
}

func Delete(log *slog.Logger, delreview DelReview) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.delete.Delete"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var id int64
	}
}
