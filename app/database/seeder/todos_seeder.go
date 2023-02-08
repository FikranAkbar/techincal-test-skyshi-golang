package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"technical-test-skyshi/app/database/entity"
)

type TodosSeeder struct {
	gormseeder.SeederAbstract
}

func NewTodosSeeder(cfg gormseeder.SeederConfiguration) *TodosSeeder {
	return &TodosSeeder{gormseeder.NewSeederAbstract(cfg)}
}

func InitTodosData() []entity.Todos {
	return []entity.Todos{
		{
			TodoId:          1,
			ActivityGroupId: 1,
			Title:           "Todo 1",
			Priority:        "normal",
		},
		{
			TodoId:          2,
			ActivityGroupId: 1,
			Title:           "Todo 2",
			Priority:        "normal",
		},
		{
			TodoId:          3,
			ActivityGroupId: 2,
			Title:           "Todo 3",
			Priority:        "normal",
		},
		{
			TodoId:          4,
			ActivityGroupId: 2,
			Title:           "Todo 4",
			Priority:        "normal",
		},
	}
}

func (s *TodosSeeder) Seed(db *gorm.DB) error {
	todos := InitTodosData()
	return db.CreateInBatches(todos, len(todos)).Error
}

func (s *TodosSeeder) Clear(db *gorm.DB) error {
	ent := entity.Todos{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
