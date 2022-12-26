package todoHandler

import (
	"net/http"
	"strconv"
	todoService "test_service/internal/services/todo_service"
	"test_service/pkg/dto"

	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"

	msgConst "test_service/pkg/common/const"
)

type HttpHandler struct {
	service todoService.TodoService
}

func NewHttpHandler(e *echo.Echo, srv todoService.TodoService) {
	handler := &HttpHandler{
		srv,
	}
	e.GET("todo-items", handler.GetAllTodo)
	e.GET("todo-items/:id", handler.GetTodoById)
	e.POST("todo-items", handler.SaveTodo)
	e.DELETE("todo-items/:id", handler.DeleteTodoById)
	e.PATCH("todo-items/:id", handler.UpdateTodoById)

}

func (h *HttpHandler) GetAllTodo(c echo.Context) error {

	qp := c.QueryParam("activity_group_id")
	page, err := strconv.Atoi(qp)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(500, dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		})
	}
	//masuk ke service
	data, err := h.service.GetAllTodo(int64(page))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: msgErrors.ErrorDataNotFound,
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Status:  msgConst.Success,
		Message: msgConst.Success,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)
}

func (h *HttpHandler) GetTodoById(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(500, dto.ResponseDTO{
			Status:  "Failed",
			Message: "Id cant be empty",
			Data:    nil,
		})
	}
	idInt, _ := strconv.Atoi(id)

	//masuk ke service
	data, err := h.service.GetTodoById(int64(idInt))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Status:  "Success",
		Message: msgConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)
}

func (h *HttpHandler) SaveTodo(c echo.Context) error {
	postDTO := dto.TodoCreateReqDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, errSave := h.service.SaveTodo(&postDTO)
	if errSave != nil {
		return c.JSON(getStatusCode(errSave), dto.ResponseDTO{
			Status:  "Failed",
			Message: errSave.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Status:  "Success",
		Message: msgConst.SaveSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) DeleteTodoById(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(500, dto.ResponseDTO{
			Status:  "Failed",
			Message: "Id cant be empty",
			Data:    nil,
		})
	}
	idInt, _ := strconv.Atoi(id)
	//masuk ke service
	err := h.service.DeleteTodoById(int64(idInt))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Not Found",
			Message: err.Error(),
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Status:  "Success",
		Message: msgConst.GetDataSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, respon)
}

func (h *HttpHandler) UpdateTodoById(c echo.Context) error {

	postDTO := dto.TodoUpdateReqDTO{}
	//byte ke json
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(500, dto.ResponseDTO{
			Status:  "Failed",
			Message: "Id cant be empty",
			Data:    nil,
		})
	}
	idInt, _ := strconv.Atoi(id)
	//masuk ke service
	data, err := h.service.UpdateTodo(int64(idInt), &postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Not Found",
			Message: err.Error(),
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Status:  "Success",
		Message: msgConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)
}
