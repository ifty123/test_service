package activityHandler

import (
	activityService "test_service/internal/services/activity_service"

	"github.com/labstack/echo"
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
