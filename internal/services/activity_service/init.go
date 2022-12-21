package activityService

import (
	activityRepository "test_service/internal/repository/activity"
)

type service struct {
	ActivityRepo activityRepository.ActivityRepository
}

func NewService(repo activityRepository.ActivityRepository) ActivityService {
	return &service{repo}
}
