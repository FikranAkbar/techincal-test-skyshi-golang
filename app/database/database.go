package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"technical-test-skyshi/app/database/entity"
	"technical-test-skyshi/config"
	"technical-test-skyshi/helper"
)

func NewDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.CreateDBURL()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.Activities{}, &entity.Todos{})
	if err != nil {
		helper.PanicIfError(err)
	}

	return db
}
