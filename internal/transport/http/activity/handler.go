package activityHandler

import (
	"net/http"
	"strconv"
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
	e.GET("activity-groups", handler.GetAllActivity)
	e.GET("activity-groups/:id", handler.GetActivityById)
	e.POST("activity-groups", handler.SaveActivity)
	e.DELETE("activity-groups/:id", handler.DeleteActivityById)
	e.PATCH("activity-groups/:id", handler.UpdateActivityById)

}

func (h *HttpHandler) GetAllActivity(c echo.Context) error {
	//masuk ke service
	data, err := h.service.GetAllActivity()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: msgErrors.ErrorDataNotFound,
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

func (h *HttpHandler) GetActivityById(c echo.Context) error {

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
	data, err := h.service.GetActivityById(int64(idInt))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: msgErrors.ErrorDataNotFound,
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
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, errSave := h.service.SaveActivity(&postDTO)
	if errSave != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(errSave), dto.ResponseDTO{
			Status:  "Failed",
			Message: err.Error(),
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

func (h *HttpHandler) DeleteActivityById(c echo.Context) error {

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
	err := h.service.DeleteActivityById(int64(idInt))
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

func (h *HttpHandler) UpdateActivityById(c echo.Context) error {

	postDTO := dto.ActivityReqDTO{}
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
	data, err := h.service.UpdateActivity(int64(idInt), &postDTO)
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
