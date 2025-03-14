package requests

type Register struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Role           string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RFToken struct {
	RefreshToken string `json:"refresh_token"`
}
