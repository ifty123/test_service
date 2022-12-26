package todoService

import (
	"fmt"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) GetTodoById(id int64) (*dto.TodoRespDTO, error) {

	data, err := s.TodoRepo.GetTodoById(id)

	if err != nil {
		return nil, fmt.Errorf("Todo with ID %d Not Found", id)
	}

	data1 := assembler.ToDTOTodo(data)
	return data1, nil
}
