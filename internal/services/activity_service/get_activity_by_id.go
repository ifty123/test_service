package activityService

import (
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) GetActivityById(id int64) (*dto.ActivityRespDTO, error) {

	data, err := s.ActivityRepo.GetActivityById(id)

	if err != nil {
		return nil, msgErrors.ErrNotFound
	}

	data1 := assembler.ToDTOActivity(data)
	return data1, nil
}
