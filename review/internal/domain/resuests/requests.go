package requests

type EmailMenor struct {
	Email string `json:"email" db:"email" validate:"required,email"`
}
