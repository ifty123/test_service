package dto

import (
	"test_service/pkg/common/validator"
	"time"
)

type ActivityRespDTO struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	Id        int64     `json:"id"`
}

type ActivityReqDTO struct {
	Email string `json:"email" validname:"email"`
	Title string `json:"title" valid:"required" validname:"title"`
}

func (dto *ActivityReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
