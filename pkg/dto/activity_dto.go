package dto

import (
	"time"
)

type ActivityRespDTO struct {
	Id        int64       `json:"id"`
	Email     string      `json:"email"`
	Title     string      `json:"title"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}
