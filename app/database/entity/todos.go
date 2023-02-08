package entity

import (
	"gorm.io/gorm"
	"time"
)

type Todos struct {
	TodoId          uint       `gorm:"column:todo_id;primaryKey"`
	ActivityGroupId uint       `gorm:"column:activity_group_id"`
	Activity        Activities `gorm:"foreignKey:ActivityGroupId;joinForeignKey:ActivityId"`
	IsActive        bool       `gorm:"column:is_active;type:boolean;default:true"`
	Title           string     `gorm:"column:title;type:varchar(300);not null"`
	Priority        string     `gorm:"column:priority;default:normal"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (Todos) TableName() string {
	return "todos"
}
