package activities_service

import (
	"context"
	"technical-test-skyshi/model/web/activity_group"
)

type ActivitiesService interface {
	GetActivities(ctx context.Context) []activity_group.ActivityGroupResponse
	GetActivityById(ctx context.Context, activityId uint) activity_group.ActivityGroupResponse
	CreateNewActivityGroup(ctx context.Context, request activity_group.ActivityGroupRequest) activity_group.ActivityGroupResponse
	DeleteActivityGroupById(ctx context.Context, activityId uint)
	UpdateActivityGroupById(ctx context.Context, activityId uint, request activity_group.ActivityGroupRequest) activity_group.ActivityGroupResponse
}
