package assembler

import (
	activity_model "test_service/internal/models/activity"
	"test_service/pkg/dto"
)

func ToDTOActivity(model *activity_model.ActivityModel) *dto.ActivityRespDTO {
	return &dto.ActivityRespDTO{
		Id:        model.Id,
		Email:     model.Email,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.Updatedat,
		DeletedAt: model.DeletedAt,
	}
}

func ToListDTOActivity(models []*activity_model.ActivityModel) []*dto.ActivityRespDTO {
	var result []*dto.ActivityRespDTO
	for _, m := range models {
		result = append(result, ToDTOActivity(m))
	}

	return result
}
