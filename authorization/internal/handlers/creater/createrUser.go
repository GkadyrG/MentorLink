package creater

import (
	"mentorlink/internal/domain/model"
)

type UserCreater interface {
	CreateUser(u *model.User) error
}

// func Refister(slog *slog.Logger, userCreater UserCreater) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		const op = "handlers.creater.createrUser.New"
// 		log := slog.With()
// 	}
// }
