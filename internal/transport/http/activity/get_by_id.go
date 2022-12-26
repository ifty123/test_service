package activityHandler

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

func (h *HttpHandler) GetActivityById(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(400, dto.ResponseDTO{
			Status:  msgErrors.ErrorFailed,
			Message: "Id cant be empty",
			Data:    nil,
		})
	}
	idInt, _ := strconv.Atoi(id)
	//masuk ke service
	data, err := h.service.GetActivityById(int64(idInt))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  msgErrors.ErrorDataNotFound,
			Message: fmt.Sprintf("Activity with ID %d Not Found", idInt),
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
