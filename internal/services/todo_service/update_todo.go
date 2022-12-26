package todoService

import (
	"fmt"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) UpdateTodo(id int64, dto *dto.TodoUpdateReqDTO) (*dto.TodoRespDTO, error) {

	_, errGet := s.TodoRepo.GetTodoById(id)
	if errGet != nil {
		return nil, fmt.Errorf("Todo with ID %d Not Found", id)
	}

	update := assembler.ToModelUpdateTodo(dto)

	err := s.TodoRepo.UpdateTodo(update)
	if err != nil {
		return nil, err
	}

	ActById, errGet := s.GetTodoById(id)
	if errGet != nil {
		return nil, err
	}

	return ActById, nil
}
