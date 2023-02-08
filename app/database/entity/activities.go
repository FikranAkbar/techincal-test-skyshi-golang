package entity

import (
	"gorm.io/gorm"
	"time"
)

type Activities struct {
	ActivityId uint   `gorm:"column:activity_id;primaryKey"`
	Title      string `gorm:"column:title;type:varchar(300);not null"`
	Email      string `gorm:"column:email;type:varchar(300)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (Activities) TableName() string {
	return "activity_group"
}
