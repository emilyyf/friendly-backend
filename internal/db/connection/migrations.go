package db

import (
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&log.Log{})
}
