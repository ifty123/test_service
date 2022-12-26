package todoHandler

import (
	"net/http"
	"strconv"
	msgConst "test_service/pkg/common/const"
	"test_service/pkg/dto"
	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

func (h *HttpHandler) GetAllTodo(c echo.Context) error {

	qp := c.QueryParam("activity_group_id")
	page, err := strconv.Atoi(qp)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(400, dto.ResponseDTO{
			Status:  msgErrors.ErrorFailed,
			Message: err.Error(),
			Data:    nil,
		})
	}
	//masuk ke service
	data, err := h.service.GetAllTodo(int64(page))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  msgErrors.ErrorFailed,
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
