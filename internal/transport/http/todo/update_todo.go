package todoHandler

import (
	"fmt"
	"net/http"
	"strconv"
	msgConst "test_service/pkg/common/const"
	"test_service/pkg/dto"
	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

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
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Status:  msgErrors.ErrorBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Status:  msgErrors.ErrorFailed,
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
			Status:  msgErrors.ErrorDataNotFound,
			Message: fmt.Sprintf("Todo with ID %d Not Found", idInt),
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
