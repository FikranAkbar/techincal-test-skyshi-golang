package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"log"
	"technical-test-skyshi/helper"
)

func getSeeders() []gormseeder.SeederInterface {
	return []gormseeder.SeederInterface{
		NewActivitiesSeeder(gormseeder.SeederConfiguration{}),
		NewTodosSeeder(gormseeder.SeederConfiguration{}),
	}
}

func SeedDB(db *gorm.DB) {
	log.Printf("Seed: Start")

	seedersStack := gormseeder.NewSeedersStack(db)
	for _, modelSeeder := range getSeeders() {
		seedersStack.AddSeeder(modelSeeder)
	}
	err := seedersStack.Seed()
	helper.PanicIfError(err)

	log.Printf("Seed: Success")
}
