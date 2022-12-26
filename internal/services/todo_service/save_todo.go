package todoService

import (
	"log"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) SaveTodo(dto *dto.TodoCreateReqDTO) (*dto.TodoRespDTO, error) {
	//assembler
	saveAct := assembler.ToModelCreateTodo(dto)

	if saveAct.Priority == "" {
		saveAct.Priority = "very-high"
	}

	id, err := s.TodoRepo.SaveTodo(saveAct)
	if err != nil {
		log.Println("err save todo :", err)
		return nil, err
	}

	ActById, errGet := s.GetTodoById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	return ActById, nil
}
