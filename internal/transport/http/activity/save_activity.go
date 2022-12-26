package activityHandler

import (
	"net/http"
	msgConst "test_service/pkg/common/const"
	"test_service/pkg/dto"

	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

func (h *HttpHandler) SaveActivity(c echo.Context) error {
	postDTO := dto.ActivityReqDTO{}
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

	data, errSave := h.service.SaveActivity(&postDTO)
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
