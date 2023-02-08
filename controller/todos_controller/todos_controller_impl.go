package todos_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/web/todo_item"
	"technical-test-skyshi/service/todos_service"
)

type TodosControllerImpl struct {
	todos_service.TodosService
}

func NewTodosControllerImpl(todosService todos_service.TodosService) *TodosControllerImpl {
	return &TodosControllerImpl{TodosService: todosService}
}

func (controller *TodosControllerImpl) GetTodos(c echo.Context) error {
	var todosResponse []todo_item.TodoItemResponse
	activityGroupId, err := strconv.Atoi(c.QueryParam("activity_group_id"))
	if err != nil {
		todosResponse = controller.TodosService.GetTodos(c.Request().Context(), nil)
	} else {
		newActivityGroupId := uint(activityGroupId)
		todosResponse = controller.TodosService.GetTodos(c.Request().Context(), &newActivityGroupId)
	}

	return helper.SuccessAPIResponse(c, todosResponse)
}

func (controller *TodosControllerImpl) GetTodoById(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("todoItemId"))
	helper.PanicIfError(err)

	todoResponse := controller.TodosService.GetTodoById(c.Request().Context(), uint(todoId))

	return helper.SuccessAPIResponse(c, todoResponse)
}

func (controller *TodosControllerImpl) CreateNewTodo(c echo.Context) error {
	var request todo_item.TodoItemRequest
	err := c.Bind(&request)
	helper.PanicIfError(err)

	todoResponse := controller.TodosService.CreateNewTodo(c.Request().Context(), request)

	return helper.SuccessAPIResponse(c, todoResponse)
}

func (controller *TodosControllerImpl) DeleteTodoById(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("todoItemId"))
	helper.PanicIfError(err)

	controller.TodosService.DeleteTodoById(c.Request().Context(), uint(todoId))

	return helper.SuccessAPIResponse(c, nil)
}

func (controller *TodosControllerImpl) UpdateTodoById(c echo.Context) error {
	var request todo_item.TodoItemRequest
	err := c.Bind(&request)
	helper.PanicIfError(err)

	todoId, err := strconv.Atoi(c.Param("todoItemId"))
	helper.PanicIfError(err)

	todoResponse := controller.TodosService.UpdateTodoById(c.Request().Context(), request, uint(todoId))

	return helper.SuccessAPIResponse(c, todoResponse)
}
