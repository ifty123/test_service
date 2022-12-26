package dto

import (
	"test_service/pkg/common/validator"
	"time"
)

type TodoRespDTO struct {
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     interface{} `json:"deleted_at"`
	Title         string      `json:"title"`
	Priority      string      `json:"priority"`
	ActiveGroupId int64       `json:"activity_group_id"`
	Id            int64       `json:"id"`
	IsActive      bool        `json:"is_active"`
}

type TodoUpdateReqDTO struct {
	Title    string `json:"title"  validname:"title"`
	Priority string `json:"priority"  validname:"priority"`
	IsActive bool   `json:"is_active"  validname:"is_active"`
	Status   string `json:"status"  validname:"status"`
}

type TodoCreateReqDTO struct {
	Title           string `json:"title"  validname:"title"`
	IsActive        bool   `json:"is_active"  validname:"is_active"`
	ActivityGroupId int64  `json:"activity_group_id"  validname:"activity_group_id"`
}

func (dto *TodoCreateReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

func (dto *TodoUpdateReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
