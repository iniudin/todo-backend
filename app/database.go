package app

import (
	"todo-backend/helper"
	"todo-backend/model/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.Todo{})
	return db
}
