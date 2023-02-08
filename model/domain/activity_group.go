package domain

import (
	"gorm.io/gorm"
	"time"
)

type ActivityGroup struct {
	ActivityId uint
	Title      string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
