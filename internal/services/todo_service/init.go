package todoService

import (
	activityRepository "test_service/internal/repository/activity"
	todoRepository "test_service/internal/repository/todo"
)

type service struct {
	TodoRepo     todoRepository.TodoRepository
	ActivityRepo activityRepository.ActivityRepository
}

func NewService(repo todoRepository.TodoRepository, activityRepo activityRepository.ActivityRepository) TodoService {
	return &service{repo, activityRepo}
}
