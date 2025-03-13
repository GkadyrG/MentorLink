package registration

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
		const op = "handlers.creater.createrUser.New"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req requests.Register

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}

		if req.Password != req.RepeatPassword {
			log.Error("passwords do not match")
			render.JSON(w, r, response.Error("passwords do not match"))
			return
		}

		if req.Role != model.RoleAdmin && req.Role != model.RoleMentor && req.Role != model.RoleUser {
			log.Error("worng with role", slog.String("role", req.Role))
			render.JSON(w, r, response.Error("wrnog with role"))
			return
		}

		existing, err := userCreater.GetByEmail(req.Email)
		if err != nil && !(errors.Is(err, db.ErrUserNotFound)) {
			log.Error("error with GetByEmail")
			render.JSON(w, r, response.Error("wtf"))
			return
		}

		if existing != nil {
			log.Error("user already exists")
			render.JSON(w, r, response.Error("user already exists"))
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			render.JSON(w, r, "error with hash")
			log.Debug("Что-то с хешированием пароля")
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: string(hash),
			Role:     req.Role,
		}

		if err = userCreater.CreateUser(user); err != nil {
			log.Error("Error with storage")
			render.JSON(w, r, response.Error("error with storage"))
			return
		}

		render.JSON(w, r, user.ID)

	}
}
