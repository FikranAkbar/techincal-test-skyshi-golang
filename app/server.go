package app

import (
	"github.com/labstack/echo/v4"
	"technical-test-skyshi/app/router"
	"technical-test-skyshi/app/router/middleware"
	"technical-test-skyshi/controller/activities_controller"
	"technical-test-skyshi/controller/todos_controller"
	"technical-test-skyshi/exception"
)

func InitServerWithEcho(activitiesController activities_controller.ActivitiesController, todosController todos_controller.TodosController) *echo.Echo {
	e := echo.New()

	router.InitRoutes(activitiesController, todosController, e)

	middleware.InitMiddleware(e)
	e.HTTPErrorHandler = exception.ErrorHandler

	return e
}
