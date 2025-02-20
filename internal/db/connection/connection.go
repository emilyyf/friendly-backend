package db

import (
	"fmt"
	"friendly-backend/internal/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost = config.GetFromEnv("POSTGRES_HOST")
	dbPort = config.GetFromEnv("POSTGRES_PORT")
	dbUser = config.GetFromEnv("POSTGRES_USER")
	dbPass = config.GetFromEnv("POSTGRES_PASSWORD")
	dbName = config.GetFromEnv("POSTGRES_DB")
	dbZone = config.GetFromEnv("POSTGRES_TIME_ZONE")
)

func OpenConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		dbZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Succesfuly connected to the database!")
	return db, nil
}
