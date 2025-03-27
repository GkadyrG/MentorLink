package model

type Mentor struct {
	ID            int64
	MentorEmail   string
	Contact       string
	CountReviews  int64
	SumRating     float32
	AverageRating float32
}
