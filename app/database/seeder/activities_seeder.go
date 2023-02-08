package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"technical-test-skyshi/app/database/entity"
)

type ActivitiesSeeder struct {
	gormseeder.SeederAbstract
}

func NewActivitiesSeeder(cfg gormseeder.SeederConfiguration) *ActivitiesSeeder {
	return &ActivitiesSeeder{gormseeder.NewSeederAbstract(cfg)}
}

func InitActivitiesData() []entity.Activities {
	return []entity.Activities{
		{
			ActivityId: 1,
			Title:      "Activity 1",
			Email:      "fikranakbar756@gmail.com",
		},
		{
			ActivityId: 2,
			Title:      "Activity 2",
			Email:      "fikranakbar756@gmail.com",
		},
	}
}

func (s *ActivitiesSeeder) Seed(db *gorm.DB) error {
	activities := InitActivitiesData()
	return db.CreateInBatches(activities, len(activities)).Error
}

func (s *ActivitiesSeeder) Clear(db *gorm.DB) error {
	ent := entity.Activities{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
