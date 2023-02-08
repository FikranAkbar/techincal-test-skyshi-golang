package activities_repository

import (
	"context"
	"gorm.io/gorm"
	"technical-test-skyshi/model/domain"
)

type ActivitiesRepository interface {
	GetActivities(ctx context.Context, db *gorm.DB) ([]domain.ActivityGroup, error)
	GetActivityById(ctx context.Context, db *gorm.DB, activityId uint) (domain.ActivityGroup, error)
	CreateNewActivityGroup(ctx context.Context, db *gorm.DB, activityGroup *domain.ActivityGroup) error
	DeleteActivityGroupById(ctx context.Context, db *gorm.DB, activityId uint) error
	UpdateActivityGroupById(ctx context.Context, db *gorm.DB, activityGroup *domain.ActivityGroup) error
}
