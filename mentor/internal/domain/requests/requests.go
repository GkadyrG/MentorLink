package requests

type RatingRequest struct {
	MentorEmail string  `json:"mentor_email"`
	Rating      float32 `json:"rating"`
}

type MentorRequest struct {
	MentorEmail string `json:"mentor_email"`
	Contact     string `json:"contact"`
}
