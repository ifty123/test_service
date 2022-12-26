package todo_model

import (
	"database/sql"
	"time"
)

type TodoModel struct {
	CreatedAt       time.Time    `db:"created_at"`
	Updatedat       time.Time    `db:"updated_at"`
	DeletedAt       sql.NullTime `db:"deleted_at"`
	Title           string       `db:"title"`
	Priority        string       `db:"priority"`
	Id              int64        `db:"id"`
	ActivityGroupId int64        `db:"activity_group_id"`
	IsActive        bool         `db:"is_active"`
}
