package todoService

import (
	"log"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) UpdateTodo(id int64, dto *dto.TodoUpdateReqDTO) (*dto.TodoRespDTO, error) {

	_, errGet := s.TodoRepo.GetTodoById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	update := assembler.ToModelUpdateTodo(dto)

	err := s.TodoRepo.UpdateTodo(update)
	if err != nil {
		log.Println("err update todo :", err)
		return nil, err
	}

	ActById, errGet := s.GetTodoById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	return ActById, nil
}
