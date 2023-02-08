package todos_service

import (
	"context"
	"technical-test-skyshi/model/web/todo_item"
)

type TodosService interface {
	GetTodos(ctx context.Context, activityGroupId *uint) []todo_item.TodoItemResponse
	GetTodoById(ctx context.Context, todoId uint) todo_item.TodoItemResponse
	CreateNewTodo(ctx context.Context, request todo_item.TodoItemRequest) todo_item.TodoItemResponse
	DeleteTodoById(ctx context.Context, todoId uint)
	UpdateTodoById(ctx context.Context, request todo_item.TodoItemRequest, todoId uint) todo_item.TodoItemResponse
}
