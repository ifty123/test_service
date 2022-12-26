package activity_model

import (
	"time"
)

type ActivityModel struct {
	CreatedAt time.Time `db:"created_at"`
	Updatedat time.Time `db:"updated_at"`
	Email     string    `db:"email"`
	Title     string    `db:"title"`
	Id        int64     `db:"id"`
}
