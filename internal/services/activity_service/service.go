package activityService

import "test_service/pkg/dto"

type ActivityService interface {
	GetAllActivity() ([]*dto.ActivityRespDTO, error)
	GetActivityById(id int64) (*dto.ActivityRespDTO, error)
	SaveActivity(dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error)
	DeleteActivityById(id int64) error
	UpdateActivity(id int64, dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error)
}
