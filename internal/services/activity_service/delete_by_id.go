package activityService

import (
	"fmt"
)

//tampil semua data
func (s *service) DeleteActivityById(id int64) error {

	_, errGet := s.ActivityRepo.GetActivityById(id)
	if errGet != nil {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}

	err := s.ActivityRepo.DeleteActivityById(id)

	if err != nil {
		return err
	}
	return nil

}
