package activityService

import (
	"log"
	"test_service/pkg/dto"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) UpdateActivity(id int64, dto *dto.ActivityReqDTO) (*dto.ActivityRespDTO, error) {

	_, errGet := s.ActivityRepo.GetActivityById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	err := s.ActivityRepo.UpdateActivity(id, dto.Title)
	if err != nil {
		log.Println("err update act :", err)
		return nil, err
	}

	ActById, errGet := s.GetActivityById(id)
	if errGet != nil {
		return nil, msgErrors.ErrNotFound
	}

	return ActById, nil
}
