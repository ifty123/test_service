package activityRepository

import (
	activity_model "test_service/internal/models/activity"
)

type ActivityRepository interface {
	GetAllActivity() ([]*activity_model.ActivityModel, error)
	GetActivityById(id int64) (*activity_model.ActivityModel, error)
	SaveActivity(payload *activity_model.ActivityModel) (int64, error)
	DeleteActivityById(id int64) error
	UpdateActivity(id int64, title string) error
}
