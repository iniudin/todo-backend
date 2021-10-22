package app

import (
	"todo-backend/helper"
	"todo-backend/model/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.Todo{})
	return db
}
