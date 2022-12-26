package activityService

import (
	"log"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) SaveActivity(dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error) {
	//assembler
	saveAct := assembler.ToModelActivity(dto)
	id, err := s.ActivityRepo.SaveActivity(saveAct)
	if err != nil {
		log.Println("err save act :", err)
		return nil, err
	}

	ActById, errGet := s.GetActivityById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	return ActById, nil
}
