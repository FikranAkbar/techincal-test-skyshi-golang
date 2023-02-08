package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	TodoId          uint
	ActivityGroupId *uint
	IsActive        *bool
	Title           *string
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
