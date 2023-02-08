package router

import (
	"github.com/labstack/echo/v4"
	"technical-test-skyshi/controller/activities_controller"
	"technical-test-skyshi/controller/todos_controller"
)

var (
	HTTPPrefix = "http://"
	HostURL    = "localhost:3030"
	//HostURLTest = "http://localhost:3031"
)

var (
	ActivitiesGroupURLPath = "/activity-groups"
	TodoItemsURLPath       = "/todo-items"
)

var (
	ActivityGroupIdPath = "/:activityGroupId"
	TodoItemIdPath      = "/:todoItemId"
)

func InitRoutes(
	activityController activities_controller.ActivitiesController,
	todosController todos_controller.TodosController,
	e *echo.Echo,
) {
	activitiesRouteGroup := e.Group(ActivitiesGroupURLPath)
	todosRouteGroup := e.Group(TodoItemsURLPath)

	// Activities API Route
	activitiesRouteGroup.GET("", activityController.GetActivities).Name = "Get all activity group"
	activitiesRouteGroup.GET(ActivityGroupIdPath, activityController.GetActivityById).Name = "Get activity group by id"
	activitiesRouteGroup.POST("", activityController.CreateNewActivityGroup).Name = "Create new activity group"
	activitiesRouteGroup.DELETE(ActivityGroupIdPath, activityController.DeleteActivityGroupById).Name = "Delete activity group by id"
	activitiesRouteGroup.PATCH(ActivityGroupIdPath, activityController.UpdateActivityGroupById).Name = "Update activity group by id"

	// Todos API Route
	todosRouteGroup.GET("", todosController.GetTodos).Name = "Get all todo_item"
	todosRouteGroup.GET(TodoItemIdPath, todosController.GetTodoById).Name = "Get todo by id"
	todosRouteGroup.POST("", todosController.CreateNewTodo).Name = "Create new todo"
	todosRouteGroup.DELETE(TodoItemIdPath, todosController.DeleteTodoById).Name = "Delete todo by id"
	todosRouteGroup.PATCH(TodoItemIdPath, todosController.UpdateTodoById).Name = "Update todo by id"
}
