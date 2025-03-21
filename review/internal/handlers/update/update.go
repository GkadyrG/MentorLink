package update

import (
	"log/slog"
	"net/http"
	"review/internal/domain/model"
	"review/internal/domain/response"
	"review/internal/lib/logger/sl"
	"review/internal/lib/validate"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type ReviewUpdate interface {
	UpdateReview(review *model.Review) error
}

func Update(log *slog.Logger, reviewUpdate ReviewUpdate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.update.Update"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req model.Review
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid request body"))
			return
		}

		if err := validate.IsValid(req); err != nil {
			log.Error("validation error", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		if err := reviewUpdate.UpdateReview(&req); err != nil {
			log.Error("failed to update review", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]any{
			"status": "updated",
		})
	}
}
