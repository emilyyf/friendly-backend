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
	db.AutoMigrate(&e.Household{})
	db.AutoMigrate(&e.Person{})
	db.AutoMigrate(&e.Child{})
	db.AutoMigrate(&e.Vaccines{})
	db.AutoMigrate(&e.VaccineDosages{})
	db.AutoMigrate(&e.PersonVaccine{})
	db.AutoMigrate(&e.Appointment{})
	db.AutoMigrate(&e.Exams{})
	db.AutoMigrate(&e.MedicalHistory{})
	db.AutoMigrate(&e.Note{})
	db.AutoMigrate(&e.Participation{})
	db.AutoMigrate(&e.Scholarship{})
	db.AutoMigrate(&e.Ultrassonography{})
}
