package exception

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"technical-test-skyshi/model/web"
)

const (
	Unauthorized        = "Unauthorized"
	NotFound            = "Not Found"
	Forbidden           = "Forbidden"
	BadRequest          = "Bad Request"
	Validation          = "Validation"
	Duplicate           = "Duplicate"
	InternalServerError = "Internal Server Error"
)

func CheckErrorContains(err error, subStr string) bool {
	if err != nil {
		return strings.Contains(strings.ToLower(err.Error()), strings.ToLower(subStr))
	}

	return false
}

func ErrorHandler(err error, c echo.Context) {
	apiResponse := web.APIResponse{
		Message: err.Error(),
	}
	var httpStatus int

	if CheckErrorContains(err, "cannot be null") {
		apiResponse.Status = BadRequest
		httpStatus = http.StatusBadRequest
	} else if CheckErrorContains(err, "not found") {
		apiResponse.Status = NotFound
		httpStatus = http.StatusNotFound
	} else {
		apiResponse.Status = InternalServerError
		httpStatus = http.StatusInternalServerError
	}

	_ = c.JSON(httpStatus, apiResponse)
}

func badRequestError(err error, c echo.Context) {
	apiResponse := web.APIResponse{
		Status:  BadRequest,
		Message: err.Error(),
		Data:    nil,
	}

	_ = c.JSON(http.StatusBadRequest, apiResponse)
}

func internalServerError(err error, c echo.Context) {
	apiResponse := web.APIResponse{
		Status:  InternalServerError,
		Message: err.Error(),
		Data:    nil,
	}

	_ = c.JSON(http.StatusInternalServerError, apiResponse)
}
