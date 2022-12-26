package todoService

import (
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) GetAllTodo(id int64) ([]*dto.TodoRespDTO, error) {
	data, err := s.TodoRepo.GetAllTodo(id)

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToListDTOTodo(data)
	return data1, nil
}
