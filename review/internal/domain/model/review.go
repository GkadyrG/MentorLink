package model

import "time"

type Review struct {
	ID          int64     `db:"id"`
	MentorEmail string    `db:"mentor_email"`
	Rating      float32   `db:"rating"`
	Comment     string    `db:"comment"`
	UserContact string    `db:"user_contact"`
	CreatedAt   time.Time `db:"created_at"`
}
