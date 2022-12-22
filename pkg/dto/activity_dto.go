package dto

import (
	"test_service/pkg/common/validator"
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

type ActivityReqDTO struct {
	Email string `json:"email" validname:"email"`
	Title string `json:"title" valid:"required" validname:"title"`
}

func (dto *ActivityReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
