package todoService

import (
	"log"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) GetAllTodo(id int64) ([]*dto.TodoRespDTO, error) {
	data, err := s.TodoRepo.GetAllTodo(id)

	if err != nil {
		log.Println("err get :", err)
		return nil, msgErrors.ErrNotFound
	}

	data1 := assembler.ToListDTOTodo(data)
	return data1, nil
}
