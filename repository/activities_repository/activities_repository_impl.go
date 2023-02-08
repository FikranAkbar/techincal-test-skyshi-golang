package activities_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"technical-test-skyshi/app/database/entity"
	"technical-test-skyshi/exception"
	"technical-test-skyshi/helper"
	"technical-test-skyshi/model/domain"
)

type ActivitiesRepositoryImpl struct {
}

func NewActivitiesRepositoryImpl() *ActivitiesRepositoryImpl {
	return &ActivitiesRepositoryImpl{}
}

func (repository *ActivitiesRepositoryImpl) GetActivities(ctx context.Context, db *gorm.DB) ([]domain.ActivityGroup, error) {
	var activityGroupEntities []entity.Activities
	err := db.WithContext(ctx).
		Find(&activityGroupEntities).
		Error

	if err != nil {
		return nil, err
	}

	var activityGroups []domain.ActivityGroup
	for _, activityGroupEntity := range activityGroupEntities {
		var activityGroup domain.ActivityGroup
		helper.ConvertActivitiesEntityToDomain(&activityGroupEntity, &activityGroup)
		activityGroups = append(activityGroups, activityGroup)
	}

	return activityGroups, nil
}

func (repository *ActivitiesRepositoryImpl) GetActivityById(ctx context.Context, db *gorm.DB, activityId uint) (domain.ActivityGroup, error) {
	var activityGroupEntity entity.Activities
	err := db.WithContext(ctx).
		Where("activity_id = ?", activityId).
		First(&activityGroupEntity).
		Error

	if exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Errorf("Activity with ID %v Not Found", activityId)
		return domain.ActivityGroup{}, logError
	}

	var activityGroup domain.ActivityGroup
	helper.ConvertActivitiesEntityToDomain(&activityGroupEntity, &activityGroup)
	return activityGroup, nil
}

func (repository *ActivitiesRepositoryImpl) CreateNewActivityGroup(ctx context.Context, db *gorm.DB, activityGroup *domain.ActivityGroup) error {
	activityGroupEntity := entity.Activities{
		Title: activityGroup.Title,
		Email: activityGroup.Email,
	}

	err := db.WithContext(ctx).
		Create(&activityGroupEntity).Error

	if err != nil {
		return err
	}

	helper.ConvertActivitiesEntityToDomain(&activityGroupEntity, activityGroup)

	return nil
}

func (repository *ActivitiesRepositoryImpl) DeleteActivityGroupById(ctx context.Context, db *gorm.DB, activityId uint) error {
	var activityGroupEntity entity.Activities

	fmt.Println("Activity ID:", activityId)

	rowsAffected := db.WithContext(ctx).Unscoped().
		Where("activity_id = ?", activityId).
		Delete(&activityGroupEntity).RowsAffected

	if rowsAffected == 0 {
		logError := fmt.Errorf("Activity with ID %v Not Found", activityId)
		return logError
	}

	return nil
}

func (repository *ActivitiesRepositoryImpl) UpdateActivityGroupById(ctx context.Context, db *gorm.DB, activityGroup *domain.ActivityGroup) error {
	activityGroupEntity := entity.Activities{
		ActivityId: activityGroup.ActivityId,
		Title:      activityGroup.Title,
		Email:      activityGroup.Email,
	}

	rowsAffected := db.WithContext(ctx).
		Model(&activityGroupEntity).
		Where("activity_id = ?", activityGroup.ActivityId).
		Updates(map[string]any{
			"title": activityGroup.Title,
		}).First(&activityGroupEntity).RowsAffected

	if rowsAffected == 0 {
		logError := fmt.Errorf("Activity with ID %v Not Found", activityGroup.ActivityId)
		return logError
	}

	helper.ConvertActivitiesEntityToDomain(&activityGroupEntity, activityGroup)

	return nil
}
