package activityService

import (
	"log"
	"test_service/pkg/dto"
	"test_service/pkg/dto/assembler"
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) GetAllActivity() ([]*dto.ActivityRespDTO, error) {
	data, err := s.ActivityRepo.GetAllActivity()

	if err != nil {
		log.Println("err get :", err)
		return nil, msgErrors.ErrNotFound
	}

	data1 := assembler.ToListDTOActivity(data)
	return data1, nil
}
