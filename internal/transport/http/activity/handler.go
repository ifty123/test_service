package activityHandler

import (
	"net/http"
	"os"
	activityService "test_service/internal/services/activity_service"
	"test_service/pkg/dto"

	msgErrors "test_service/pkg/errors"

	"github.com/apex/log"
	"github.com/labstack/echo"

	msgConst "test_service/pkg/common/const"
)

type HttpHandler struct {
	service activityService.ActivityService
}

func NewHttpHandler(e *echo.Echo, srv activityService.ActivityService) {
	handler := &HttpHandler{
		srv,
	}
	e.GET("api/v1/latihan/ping", handler.Ping)
	e.GET("api/v1/activity-groups", handler.GetAllActivity)

}

func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

func (h *HttpHandler) GetAllActivity(c echo.Context) error {
	//masuk ke service
	data, err := h.service.GetAllActivity()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: msgErrors.ErrorDataNotFound,
			Data:    nil,
		})
	}

	respon := dto.ResponseDTO{
		Success: true,
		Message: msgConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, respon)
}
