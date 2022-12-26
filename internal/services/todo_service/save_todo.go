package todoService

import (
	"fmt"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) SaveTodo(dto *dto.TodoCreateReqDTO) (*dto.TodoRespDTO, error) {
	//assembler
	saveAct := assembler.ToModelCreateTodo(dto)

	if saveAct.Priority == "" {
		saveAct.Priority = "very-high"
	}

	//cek activity id, jika ada maka ditolak
	activity, erAct := s.ActivityRepo.GetActivityById(dto.ActivityGroupId)
	if erAct != nil || activity == nil {
		return nil, fmt.Errorf("Todo with activity ID %d Not Found", saveAct.ActivityGroupId)
	}

	id, err := s.TodoRepo.SaveTodo(saveAct)
	if err != nil {
		return nil, err
	}

	ActById, errGet := s.GetTodoById(id)
	if errGet != nil {
		return nil, err
	}

	return ActById, nil
}
