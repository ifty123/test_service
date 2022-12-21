package activity_model

import "time"

type ActivityModel struct {
	Id        int64     `db:"id"`
	Email     string    `db:"email"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	Updatedat time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
