package todoService

import (
	"fmt"
)

//tampil semua data
func (s *service) DeleteTodoById(id int64) error {

	_, errGet := s.TodoRepo.GetTodoById(id)
	if errGet != nil {
		return fmt.Errorf("Activity with ID %d Not Found", id)
	}

	err := s.TodoRepo.DeleteTodoById(id)

	if err != nil {
		return err
	}
	return nil

}
