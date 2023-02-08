package helper

import (
	"technical-test-skyshi/app/database/entity"
	"technical-test-skyshi/model/domain"
)

func ConvertActivitiesEntityToDomain(entity *entity.Activities, domain *domain.ActivityGroup) {
	domain.ActivityId = entity.ActivityId
	domain.Title = entity.Title
	domain.Email = entity.Email
	domain.CreatedAt = entity.CreatedAt
	domain.UpdatedAt = entity.UpdatedAt
	domain.DeletedAt = entity.DeletedAt
}

func ConvertTodosEntityToDomain(entity *entity.Todos, domain *domain.Todo) {
	newActivityGroupId := entity.ActivityGroupId
	newTitle := entity.Title
	newIsActive := entity.IsActive

	domain.TodoId = entity.TodoId
	domain.ActivityGroupId = &newActivityGroupId
	domain.Title = &newTitle
	domain.IsActive = &newIsActive
	domain.Priority = entity.Priority
	domain.CreatedAt = entity.CreatedAt
	domain.UpdatedAt = entity.UpdatedAt
	domain.DeletedAt = entity.DeletedAt
}
