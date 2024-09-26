package db

import (
	"database/sql"
	"embed"
	"friendly-backend/internal/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	dbDriver = config.GetFromEnv("POSTGRES_DRIVER")
	dbUrl    = config.GetFromEnv("POSTGRES_SOURCE")
)

//go:embed migrations/*.sql
var Migrations embed.FS

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Succesfuly connected to the database!")
	return db, nil
}
