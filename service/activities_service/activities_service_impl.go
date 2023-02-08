package activities_service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/domain"
	"technical-test-skyshi/model/web/activity_group"
	"technical-test-skyshi/repository/activities_repository"
)

type ActivitiesServiceImpl struct {
	activities_repository.ActivitiesRepository
	*gorm.DB
}

func NewActivitiesServiceImpl(activitiesRepository activities_repository.ActivitiesRepository, DB *gorm.DB) *ActivitiesServiceImpl {
	return &ActivitiesServiceImpl{ActivitiesRepository: activitiesRepository, DB: DB}
}

func (service *ActivitiesServiceImpl) GetActivities(ctx context.Context) []activity_group.ActivityGroupResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	activityGroups, err := service.ActivitiesRepository.GetActivities(ctx, tx)
	helper.PanicIfError(err)

	var activityGroupResponses []activity_group.ActivityGroupResponse
	for _, activityGroup := range activityGroups {
		activityGroupResponses = append(activityGroupResponses, activity_group.ActivityGroupResponse{
			ActivityId: activityGroup.ActivityId,
			Email:      activityGroup.Email,
			Title:      activityGroup.Title,
			CreatedAt:  activityGroup.CreatedAt,
			UpdatedAt:  activityGroup.UpdatedAt,
			DeletedAt:  activityGroup.DeletedAt,
		})
	}

	return activityGroupResponses
}

func (service *ActivitiesServiceImpl) GetActivityById(ctx context.Context, activityId uint) activity_group.ActivityGroupResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	activityGroup, err := service.ActivitiesRepository.GetActivityById(ctx, tx, activityId)
	helper.PanicIfError(err)

	return activity_group.ActivityGroupResponse{
		ActivityId: activityGroup.ActivityId,
		Email:      activityGroup.Email,
		Title:      activityGroup.Title,
		CreatedAt:  activityGroup.CreatedAt,
		UpdatedAt:  activityGroup.UpdatedAt,
		DeletedAt:  activityGroup.DeletedAt,
	}
}

func (service *ActivitiesServiceImpl) CreateNewActivityGroup(ctx context.Context, request activity_group.ActivityGroupRequest) activity_group.ActivityGroupResponse {
	if request.Title == "" {
		helper.PanicIfError(errors.New("title cannot be null"))
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	activityGroup := domain.ActivityGroup{Email: request.Email, Title: request.Title}
	err := service.ActivitiesRepository.CreateNewActivityGroup(ctx, tx, &activityGroup)
	helper.PanicIfError(err)

	return activity_group.ActivityGroupResponse{
		ActivityId: activityGroup.ActivityId,
		Email:      activityGroup.Email,
		Title:      activityGroup.Title,
		CreatedAt:  activityGroup.CreatedAt,
		UpdatedAt:  activityGroup.UpdatedAt,
		DeletedAt:  activityGroup.DeletedAt,
	}
}

func (service *ActivitiesServiceImpl) DeleteActivityGroupById(ctx context.Context, activityId uint) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	err := service.ActivitiesRepository.DeleteActivityGroupById(ctx, tx, activityId)
	helper.PanicIfError(err)
}

func (service *ActivitiesServiceImpl) UpdateActivityGroupById(ctx context.Context, activityId uint, request activity_group.ActivityGroupRequest) activity_group.ActivityGroupResponse {
	if request.Title == "" {
		helper.PanicIfError(errors.New("title cannot be null"))
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	activityGroup := &domain.ActivityGroup{ActivityId: activityId, Email: request.Email, Title: request.Title}
	err := service.ActivitiesRepository.UpdateActivityGroupById(ctx, tx, activityGroup)
	helper.PanicIfError(err)

	return activity_group.ActivityGroupResponse{
		ActivityId: activityGroup.ActivityId,
		Email:      activityGroup.Email,
		Title:      activityGroup.Title,
		CreatedAt:  activityGroup.CreatedAt,
		UpdatedAt:  activityGroup.UpdatedAt,
		DeletedAt:  activityGroup.DeletedAt,
	}
}
