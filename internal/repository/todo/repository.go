package todoRepository

import todo_model "test_service/internal/models/todo"

type TodoRepository interface {
	GetAllTodo(id int64) ([]*todo_model.TodoModel, error)
	GetTodoById(id int64) (*todo_model.TodoModel, error)
	SaveTodo(payload *todo_model.TodoModel) (int64, error)
	DeleteTodoById(id int64) error
	UpdateTodo(payload *todo_model.TodoModel) error
}
