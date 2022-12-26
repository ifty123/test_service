package todo_model

import (
	"time"
)

type TodoModel struct {
	CreatedAt       time.Time `db:"created_at"`
	Updatedat       time.Time `db:"updated_at"`
	Title           string    `db:"title"`
	Priority        string    `db:"priority"`
	Id              int64     `db:"id"`
	ActivityGroupId int64     `db:"activity_group_id"`
	IsActive        bool      `db:"is_active"`
}
