package db

import (
	e "friendly-backend/internal/db/entities"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&e.User{})
	db.AutoMigrate(&e.Log{})
}
