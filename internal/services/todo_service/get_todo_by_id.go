package todoService

import (
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) GetTodoById(id int64) (*dto.TodoRespDTO, error) {

	data, err := s.TodoRepo.GetTodoById(id)

	if err != nil {
		return nil, msgErrors.ErrNotFound
	}

	data1 := assembler.ToDTOTodo(data)
	return data1, nil
}
