package todos_repository

import (
	"context"
	"gorm.io/gorm"
	"technical-test-skyshi/model/domain"
)

type TodosRepository interface {
	GetTodos(ctx context.Context, db *gorm.DB, activityGroupId *uint) ([]domain.Todo, error)
	GetTodoById(ctx context.Context, db *gorm.DB, todoId uint) (domain.Todo, error)
	CreateNewTodo(ctx context.Context, db *gorm.DB, todoItem *domain.Todo) error
	DeleteTodoById(ctx context.Context, db *gorm.DB, todoId uint) error
	UpdateTodoById(ctx context.Context, db *gorm.DB, todoItem *domain.Todo) error
}
