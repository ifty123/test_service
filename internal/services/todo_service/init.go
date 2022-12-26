package todoService

import (
	todoRepository "test_service/internal/repository/todo"
)

type service struct {
	TodoRepo todoRepository.TodoRepository
}

func NewService(repo todoRepository.TodoRepository) TodoService {
	return &service{repo}
}
