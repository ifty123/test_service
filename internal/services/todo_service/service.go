package todoService

import "test_service/pkg/dto"

type TodoService interface {
	GetAllTodo(id int64) ([]*dto.TodoRespDTO, error)
	GetTodoById(id int64) (*dto.TodoRespDTO, error)
	SaveTodo(dto *dto.TodoCreateReqDTO) (*dto.TodoRespDTO, error)
	DeleteTodoById(id int64) error
	UpdateTodo(id int64, dto *dto.TodoUpdateReqDTO) (*dto.TodoRespDTO, error)
}
