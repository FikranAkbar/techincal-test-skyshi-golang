package todos_controller

import "github.com/labstack/echo/v4"

type TodosController interface {
	GetTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateNewTodo(c echo.Context) error
	DeleteTodoById(c echo.Context) error
	UpdateTodoById(c echo.Context) error
}
