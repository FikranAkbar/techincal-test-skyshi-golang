package activity_group

import (
	"gorm.io/gorm"
	"time"
)

type ActivityGroupResponse struct {
	ActivityId uint           `json:"activity_id" mapstructure:"activity_id"`
	Email      string         `json:"email" mapstructure:"email"`
	Title      string         `json:"title" mapstructure:"title"`
	CreatedAt  time.Time      `json:"created_at" mapstructure:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at"`
}
