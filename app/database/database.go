package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"technical-test-skyshi/config"
	"technical-test-skyshi/helper"
)

func NewDB(config *config.Config) *gorm.DB {
	//db, err := gorm.Open(mysql.Open(config.CreateDBURL()), &gorm.Config{
	//	SkipDefaultTransaction: true,
	//})
	//helper.PanicIfError(err)

	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"password",
		"localhost",
		"3306",
		"technical_test_skyshi_database",
	)

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}
