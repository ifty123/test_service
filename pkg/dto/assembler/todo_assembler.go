package assembler

import (
	todo_model "test_service/internal/models/todo"
	"test_service/pkg/dto"
)

func ToDTOTodo(model *todo_model.TodoModel) *dto.TodoRespDTO {
	var time interface{}

	if model.DeletedAt.Valid {
		time = model.DeletedAt.Time
	} else {
		time = nil
	}

	return &dto.TodoRespDTO{
		Id:            model.Id,
		Title:         model.Title,
		ActiveGroupId: model.ActivityGroupId,
		IsActive:      model.IsActive,
		Priority:      model.Priority,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.Updatedat,
		DeletedAt:     time,
	}
}

func ToListDTOTodo(models []*todo_model.TodoModel) []*dto.TodoRespDTO {
	var result []*dto.TodoRespDTO
	for _, m := range models {
		result = append(result, ToDTOTodo(m))
	}

	return result
}

func ToModelCreateTodo(dto *dto.TodoCreateReqDTO) *todo_model.TodoModel {

	return &todo_model.TodoModel{
		Title:           dto.Title,
		ActivityGroupId: dto.ActivityGroupId,
		IsActive:        dto.IsActive,
	}
}

func ToModelUpdateTodo(dto *dto.TodoUpdateReqDTO) *todo_model.TodoModel {

	return &todo_model.TodoModel{
		Title:    dto.Title,
		Priority: dto.Priority,
		IsActive: dto.IsActive,
	}
}
