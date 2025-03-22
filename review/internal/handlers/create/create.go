package create

import (
	"log/slog"
	"net/http"
	"review/internal/domain/model"
	"review/internal/domain/response"
	"review/internal/lib/logger/sl"
	"review/internal/lib/validate"
	mwAuth "review/internal/middleware/auth"
	"review/pkg/token"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type ReviewCreater interface {
	CreateReview(review *model.Review) (int64, error)
}

func Create(log *slog.Logger, reviewCreater ReviewCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.create.Create"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		claims := r.Context().Value(mwAuth.UserKey).(*token.Claims)

		var req model.Review
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid request body"))
			return
		}

		req.UserID = claims.UserID

		if err := validate.IsValid(req); err != nil {
			log.Error("validation error", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		req.CreatedAt = time.Now()
		id, err := reviewCreater.CreateReview(&req)
		if err != nil {
			log.Error("falied to create review", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error("server error"))
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]any{
			"id": id,
		})
	}
}
