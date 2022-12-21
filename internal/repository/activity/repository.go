package activityRepository

import (
	activity_model "test_service/internal/models/activity"
)

type ActivityRepository interface {
	GetAllActivity() ([]*activity_model.ActivityModel, error)
}
