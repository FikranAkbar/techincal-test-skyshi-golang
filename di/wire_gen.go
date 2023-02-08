// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"technical-test-skyshi/app"
	"technical-test-skyshi/app/database"
	"technical-test-skyshi/config"
	"technical-test-skyshi/controller/activities_controller"
	"technical-test-skyshi/controller/todos_controller"
	"technical-test-skyshi/repository/activities_repository"
	"technical-test-skyshi/repository/todos_repository"
	"technical-test-skyshi/service/activities_service"
	"technical-test-skyshi/service/todos_service"
)

// Injectors from injector.go:

func InitializedEchoServer() *echo.Echo {
	activitiesRepositoryImpl := activities_repository.NewActivitiesRepositoryImpl()
	configConfig := config.LoadConfigFromEnv()
	db := database.NewDB(configConfig)
	activitiesServiceImpl := activities_service.NewActivitiesServiceImpl(activitiesRepositoryImpl, db)
	activitiesControllerImpl := activities_controller.NewActivitiesControllerImpl(activitiesServiceImpl)
	todosRepositoryImpl := todos_repository.NewTodosRepositoryImpl()
	todosServiceImpl := todos_service.NewTodosServiceImpl(todosRepositoryImpl, db)
	todosControllerImpl := todos_controller.NewTodosControllerImpl(todosServiceImpl)
	echoEcho := app.InitServerWithEcho(activitiesControllerImpl, todosControllerImpl)
	return echoEcho
}

// injector.go:

var activitiesSet = wire.NewSet(activities_repository.NewActivitiesRepositoryImpl, wire.Bind(new(activities_repository.ActivitiesRepository), new(*activities_repository.ActivitiesRepositoryImpl)), activities_service.NewActivitiesServiceImpl, wire.Bind(new(activities_service.ActivitiesService), new(*activities_service.ActivitiesServiceImpl)), activities_controller.NewActivitiesControllerImpl, wire.Bind(new(activities_controller.ActivitiesController), new(*activities_controller.ActivitiesControllerImpl)))

var todosSet = wire.NewSet(todos_repository.NewTodosRepositoryImpl, wire.Bind(new(todos_repository.TodosRepository), new(*todos_repository.TodosRepositoryImpl)), todos_service.NewTodosServiceImpl, wire.Bind(new(todos_service.TodosService), new(*todos_service.TodosServiceImpl)), todos_controller.NewTodosControllerImpl, wire.Bind(new(todos_controller.TodosController), new(*todos_controller.TodosControllerImpl)))

var completeSet = wire.NewSet(app.InitServerWithEcho, validator.New, config.LoadConfigFromEnv, activitiesSet,
	todosSet,
)
