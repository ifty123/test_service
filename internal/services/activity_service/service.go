package activityService

import "test_service/pkg/dto"

type ActivityService interface {
	GetAllActivity() ([]*dto.ActivityRespDTO, error)
}
