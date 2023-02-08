package todo_item

import (
	"gorm.io/gorm"
	"time"
)

type TodoItemResponse struct {
	Id              uint           `json:"id" mapstructure:"id"`
	ActivityGroupId *uint          `json:"activity_group_id" mapstructure:"activity_group_id"`
	Title           *string        `json:"title" mapstructure:"title"`
	IsActive        *bool          `json:"is_active" mapstructure:"is_active"`
	Priority        string         `json:"priority" mapstructure:"priority"`
	CreatedAt       time.Time      `json:"created_at" mapstructure:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at"`
}
