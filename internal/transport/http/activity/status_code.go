package activityHandler

import (
	"net/http"
	msgErrors "test_service/pkg/errors"
)

//status error
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	//error
	switch err {
	case msgErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case msgErrors.ErrNotFound:
		return http.StatusNotFound
	case msgErrors.ErrConflict:
		return http.StatusConflict
	case msgErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case msgErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
