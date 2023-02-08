package database

import (
	"gorm.io/gorm"
	"log"
	. "technical-test-skyshi/app/database/entity"
	"technical-test-skyshi/helper"
)

func GetEntities() []any {
	return []any{
		&Activities{},
		&Todos{},
	}
}

func Migrate(db *gorm.DB) {
	log.Printf("Migrate: Start")

	entities := GetEntities()
	migrator := db.Migrator()

	for _, entity := range helper.ReverseSlice(entities) {
		dropTableErr := migrator.DropTable(entity)
		helper.PanicIfError(dropTableErr)
	}

	for _, entity := range entities {
		migrateTableErr := migrator.AutoMigrate(entity)
		helper.PanicIfError(migrateTableErr)
	}

	log.Printf("Migrate: Success")
}
