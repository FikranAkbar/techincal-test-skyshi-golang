package activities_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/web"
	"technical-test-skyshi/model/web/activity_group"
	"technical-test-skyshi/service/activities_service"
)

type ActivitiesControllerImpl struct {
	activities_service.ActivitiesService
}

func NewActivitiesControllerImpl(activitiesService activities_service.ActivitiesService) *ActivitiesControllerImpl {
	return &ActivitiesControllerImpl{ActivitiesService: activitiesService}
}

func (controller *ActivitiesControllerImpl) GetActivities(c echo.Context) error {
	activitiesResponse := controller.ActivitiesService.GetActivities(c.Request().Context())

	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activitiesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ActivitiesControllerImpl) GetActivityById(c echo.Context) error {
	activityGroupId, activityIdErr := strconv.Atoi(c.Param("activityGroupId"))
	helper.PanicIfError(activityIdErr)

	activityResponse := controller.ActivitiesService.GetActivityById(c.Request().Context(), uint(activityGroupId))

	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ActivitiesControllerImpl) CreateNewActivityGroup(c echo.Context) error {
	var request activity_group.ActivityGroupRequest
	err := c.Bind(&request)
	helper.PanicIfError(err)

	activityResponse := controller.ActivitiesService.CreateNewActivityGroup(
		c.Request().Context(), request)

	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ActivitiesControllerImpl) DeleteActivityGroupById(c echo.Context) error {
	activityId, activityIdErr := strconv.Atoi(c.Param("activityGroupId"))
	helper.PanicIfError(activityIdErr)

	controller.ActivitiesService.DeleteActivityGroupById(c.Request().Context(), uint(activityId))

	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ActivitiesControllerImpl) UpdateActivityGroupById(c echo.Context) error {
	activityId, activityIdErr := strconv.Atoi(c.Param("activityGroupId"))
	helper.PanicIfError(activityIdErr)

	var request activity_group.ActivityGroupRequest
	err := c.Bind(&request)
	helper.PanicIfError(err)

	activityResponse := controller.ActivitiesService.UpdateActivityGroupById(
		c.Request().Context(), uint(activityId), request)

	apiResponse := web.APIResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
