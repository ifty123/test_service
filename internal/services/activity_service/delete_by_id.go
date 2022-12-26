package activityService

import (
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) DeleteActivityById(id int64) error {

	_, errGet := s.ActivityRepo.GetActivityById(id)
	if errGet != nil {
		return msgErrors.ErrNotFound
	}

	err := s.ActivityRepo.DeleteActivityById(id)

	if err != nil {
		return err
	}
	return nil

}
