package todos_service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/domain"
	"technical-test-skyshi/model/web/todo_item"
	"technical-test-skyshi/repository/todos_repository"
)

type TodosServiceImpl struct {
	todos_repository.TodosRepository
	*gorm.DB
}

func NewTodosServiceImpl(todosRepository todos_repository.TodosRepository, DB *gorm.DB) *TodosServiceImpl {
	return &TodosServiceImpl{TodosRepository: todosRepository, DB: DB}
}

func (service *TodosServiceImpl) GetTodos(ctx context.Context, activityGroupId *uint) []todo_item.TodoItemResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	todos, err := service.TodosRepository.GetTodos(ctx, tx, activityGroupId)
	helper.PanicIfError(err)

	var todosResponse []todo_item.TodoItemResponse
	for _, todo := range todos {
		todosResponse = append(todosResponse, todo_item.TodoItemResponse{
			Id:              todo.TodoId,
			ActivityGroupId: todo.ActivityGroupId,
			Title:           todo.Title,
			IsActive:        todo.IsActive,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdatedAt:       todo.UpdatedAt,
			DeletedAt:       todo.DeletedAt,
		})
	}

	return todosResponse
}

func (service *TodosServiceImpl) GetTodoById(ctx context.Context, todoId uint) todo_item.TodoItemResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodosRepository.GetTodoById(ctx, tx, todoId)
	helper.PanicIfError(err)

	return todo_item.TodoItemResponse{
		Id:              todo.TodoId,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       todo.DeletedAt,
	}
}

func (service *TodosServiceImpl) CreateNewTodo(ctx context.Context, request todo_item.TodoItemRequest) todo_item.TodoItemResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if *request.Title == "" {
		helper.PanicIfError(errors.New("title cannot be null"))
	}

	if request.ActivityGroupId == nil {
		helper.PanicIfError(errors.New("activity_group_id cannot be null"))
	}

	todo := domain.Todo{
		Title:           request.Title,
		ActivityGroupId: request.ActivityGroupId,
	}

	err := service.TodosRepository.CreateNewTodo(ctx, tx, &todo)
	helper.PanicIfError(err)

	return todo_item.TodoItemResponse{
		Id:              todo.TodoId,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       todo.DeletedAt,
	}
}

func (service *TodosServiceImpl) DeleteTodoById(ctx context.Context, todoId uint) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	err := service.TodosRepository.DeleteTodoById(ctx, tx, todoId)
	helper.PanicIfError(err)
}

func (service *TodosServiceImpl) UpdateTodoById(ctx context.Context, request todo_item.TodoItemRequest, todoId uint) todo_item.TodoItemResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		TodoId:   todoId,
		Title:    request.Title,
		IsActive: request.IsActive,
	}

	err := service.TodosRepository.UpdateTodoById(ctx, tx, &todo)
	helper.PanicIfError(err)

	return todo_item.TodoItemResponse{
		Id:              todo.TodoId,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       todo.DeletedAt,
	}
}
