package assembler

import (
	activity_model "test_service/internal/models/activity"
	"test_service/pkg/dto"
)

func ToDTOActivity(model *activity_model.ActivityModel) *dto.ActivityRespDTO {
	var time interface{}
	if model.DeletedAt.Valid {
		time = model.DeletedAt.Time
	} else {
		time = nil
	}
	return &dto.ActivityRespDTO{
		Id:        model.Id,
		Email:     model.Email,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.Updatedat,
		DeletedAt: time,
	}
}

func ToListDTOActivity(models []*activity_model.ActivityModel) []*dto.ActivityRespDTO {
	var result []*dto.ActivityRespDTO
	for _, m := range models {
		result = append(result, ToDTOActivity(m))
	}

	return result
}

func ToModelActivity(dto *dto.ActivityReqDTO) *activity_model.ActivityModel {

	return &activity_model.ActivityModel{
		Email: dto.Email,
		Title: dto.Title,
	}
}
