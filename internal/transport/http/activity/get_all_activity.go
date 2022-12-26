package activityHandler

import (
	"net/http"
	msgConst "test_service/pkg/common/const"
	"test_service/pkg/dto"
	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

func (h *HttpHandler) GetAllActivity(c echo.Context) error {
	//masuk ke service
	data, err := h.service.GetAllActivity()
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
