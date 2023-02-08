package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"technical-test-skyshi/model/web"
)

func SuccessAPIResponse(c echo.Context, data any) error {
	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
