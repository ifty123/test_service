package activityService

import (
	"fmt"
	"test_service/pkg/dto"
)

//tampil semua data
func (s *service) UpdateActivity(id int64, dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error) {

	_, errGet := s.ActivityRepo.GetActivityById(id)
	if errGet != nil {
		return nil, fmt.Errorf("Activity with ID %d Not Found", id)
	}

	err := s.ActivityRepo.UpdateActivity(id, dto.Title)
	if err != nil {
		return nil, err
	}

	ActById, errGet := s.GetActivityById(id)
	if errGet != nil {
		return nil, err
	}

	return ActById, nil
}
