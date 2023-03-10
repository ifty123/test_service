package todoService

import (
	msgErrors "test_service/pkg/errors"
)

//tampil semua data
func (s *service) DeleteTodoById(id int64) error {

	_, errGet := s.TodoRepo.GetTodoById(id)
	if errGet != nil {
		return msgErrors.ErrNotFound
	}

	err := s.TodoRepo.DeleteTodoById(id)

	if err != nil {
		return err
	}
	return nil

}
