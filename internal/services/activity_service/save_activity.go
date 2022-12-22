package activityService

import (
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) SaveActivity(dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error) {
	//assembler
	saveAct := assembler.ToModelActivity(dto)
	id, err := s.ActivityRepo.SaveActivity(saveAct)
	if err != nil {
		return nil, err
	}

	ActById, errGet := s.GetActivityById(id)
	if errGet != nil {
		return nil, err
	}

	return ActById, nil
}
