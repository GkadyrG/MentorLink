package handlers

import (
	"errors"
	"log/slog"
	"mentorlink/internal/domain/model"
	"mentorlink/internal/domain/requests"
	"mentorlink/internal/domain/response"
	"mentorlink/internal/lib/logger/sl"
	"mentorlink/internal/storage/db"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserCreater interface {
	CreateUser(u *model.User) error
	GetByEmail(email string) (*model.User, error)
}

func Register(log *slog.Logger, userCreater UserCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.register.Register"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req requests.Register
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid request body"))
			return
		}

		if req.Password != req.RepeatPassword {
			log.Warn("passwords mismatch", slog.String("email", req.Email))
			render.Status(r, http.StatusBadGateway)
			render.JSON(w, r, response.Error("passwords do not match"))
			return
		}

		if !isValideRole(req.Role) {
			log.Warn("invalid role provided", slog.String("role", req.Role))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid role"))
			return
		}

		existing, err := userCreater.GetByEmail(req.Email)
		if err != nil && !errors.Is(err, db.ErrUserNotFound) {
			log.Error("failed to check user existence", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error("internal error"))
			return
		}

		if existing != nil {
			log.Error("user already exists", slog.String("email", req.Email))
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, response.Error("user already exists"))
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("failed to hash password", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "failed to process password")
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: string(hash),
			Role:     req.Role,
		}

		if err = userCreater.CreateUser(user); err != nil {
			log.Error("failed to create user", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to create user"))
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]any{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		})

	}
}

func isValideRole(role string) bool {
	switch role {
	case model.RoleAdmin, model.RoleMentor, model.RoleUser:
		return true
	default:
		return false
	}
}
