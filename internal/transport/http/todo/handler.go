package todoHandler

import (
	activityService "test_service/internal/services/activity_service"
	todoService "test_service/internal/services/todo_service"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	service    todoService.TodoService
	actService activityService.ActivityService
}

func NewHttpHandler(e *echo.Echo, srv todoService.TodoService, act activityService.ActivityService) {
	handler := &HttpHandler{
		srv,
		act,
	}
	e.GET("todo-items", handler.GetAllTodo)
	e.GET("todo-items/:id", handler.GetTodoById)
	e.POST("todo-items", handler.SaveTodo)
	e.DELETE("todo-items/:id", handler.DeleteTodoById)
	e.PATCH("todo-items/:id", handler.UpdateTodoById)

}
