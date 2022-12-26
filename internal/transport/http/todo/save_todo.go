package todoHandler

import (
	"fmt"
	"net/http"
	msgConst "test_service/pkg/common/const"
	"test_service/pkg/dto"
	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

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
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Status:  msgErrors.ErrorBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	act, errAct := h.actService.GetActivityById(postDTO.ActivityGroupId)
	if errAct != nil || act == nil {
		log.Error(errAct.Error())
		return c.JSON(getStatusCode(errAct), dto.ResponseDTO{
			Status:  msgErrors.ErrorDataNotFound,
			Message: fmt.Sprintf("Activity with ID %d Not Found", postDTO.ActivityGroupId),
			Data:    nil,
		})
	}

	data, errSave := h.service.SaveTodo(&postDTO)
	if errSave != nil {
		return c.JSON(getStatusCode(errSave), dto.ResponseDTO{
			Status:  msgErrors.ErrorBadRequest,
			Message: errSave.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Status:  msgConst.Success,
		Message: msgConst.Success,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}
