package todos_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"technical-test-skyshi/app/database/entity"
	"technical-test-skyshi/exception"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/domain"
)

type TodosRepositoryImpl struct {
}

func NewTodosRepositoryImpl() *TodosRepositoryImpl {
	return &TodosRepositoryImpl{}
}

func (repository *TodosRepositoryImpl) GetTodos(ctx context.Context, db *gorm.DB, activityGroupId *uint) ([]domain.Todo, error) {
	var todoEntities []entity.Todos
	var err error

	if activityGroupId == nil {
		err = db.WithContext(ctx).
			Find(&todoEntities).Error
	} else {
		err = db.WithContext(ctx).
			Where("activity_group_id = ?", activityGroupId).
			Find(&todoEntities).Error
	}

	if err != nil {
		return nil, err
	}

	var todoItems []domain.Todo
	for _, todoEntity := range todoEntities {
		todoItem := domain.Todo{}
		helper.ConvertTodosEntityToDomain(&todoEntity, &todoItem)
		todoItems = append(todoItems, todoItem)
	}

	return todoItems, nil
}

func (repository *TodosRepositoryImpl) GetTodoById(ctx context.Context, db *gorm.DB, todoId uint) (domain.Todo, error) {
	var todoEntity entity.Todos
	err := db.WithContext(ctx).
		Where("todo_id = ?", todoId).
		First(&todoEntity).Error

	if exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Errorf("Todo with ID %v Not Found", todoId)
		return domain.Todo{}, logError
	}

	var todoItem domain.Todo
	helper.ConvertTodosEntityToDomain(&todoEntity, &todoItem)
	return todoItem, nil
}

func (repository *TodosRepositoryImpl) CreateNewTodo(ctx context.Context, db *gorm.DB, todoItem *domain.Todo) error {
	todoEntity := entity.Todos{
		Title:           *todoItem.Title,
		ActivityGroupId: *todoItem.ActivityGroupId,
	}

	var activityGroupEntity entity.Activities
	err := db.WithContext(ctx).
		Where("activity_id = ?", todoItem.ActivityGroupId).
		First(&activityGroupEntity).Error
	if exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Errorf("Activity with ID %v Not Exist", *todoItem.ActivityGroupId)
		return logError
	}

	err = db.WithContext(ctx).
		Create(&todoEntity).Error
	if err != nil {
		return err
	}

	helper.ConvertTodosEntityToDomain(&todoEntity, todoItem)

	return nil
}

func (repository *TodosRepositoryImpl) DeleteTodoById(ctx context.Context, db *gorm.DB, todoId uint) error {
	var todoEntity entity.Todos

	rowsAffected := db.WithContext(ctx).
		Where("todo_id = ?", todoId).
		Delete(&todoEntity).RowsAffected

	if rowsAffected == 0 {
		logError := fmt.Errorf("Todo with ID %v Not Found", todoId)
		return logError
	}

	return nil
}

func (repository *TodosRepositoryImpl) UpdateTodoById(ctx context.Context, db *gorm.DB, todoItem *domain.Todo) error {
	var todoEntity entity.Todos
	updatedColumn := map[string]any{}

	if todoItem.Title != nil && *todoItem.Title != "" {
		updatedColumn["title"] = *todoItem.Title
	}

	if todoItem.IsActive != nil {
		updatedColumn["is_active"] = *todoItem.IsActive
	}

	rowsAffected := db.WithContext(ctx).
		Model(&todoEntity).
		Where("todo_id = ?", todoItem.TodoId).
		Updates(updatedColumn).
		First(&todoEntity).RowsAffected

	if rowsAffected == 0 {
		logError := fmt.Errorf("Todo with ID %v Not Found", todoItem.TodoId)
		return logError
	}

	helper.ConvertTodosEntityToDomain(&todoEntity, todoItem)

	return nil
}
