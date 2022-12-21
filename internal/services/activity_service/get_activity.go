package activityService

import (
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
)

//tampil semua data
func (s *service) GetAllActivity() ([]*dto.ActivityRespDTO, error) {
	data, err := s.ActivityRepo.GetAllActivity()

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToListDTOActivity(data)
	return data1, nil
}
