package activities_controller

import "github.com/labstack/echo/v4"

type ActivitiesController interface {
	GetActivities(c echo.Context) error
	GetActivityById(c echo.Context) error
	CreateNewActivityGroup(c echo.Context) error
	DeleteActivityGroupById(c echo.Context) error
	UpdateActivityGroupById(c echo.Context) error
}
