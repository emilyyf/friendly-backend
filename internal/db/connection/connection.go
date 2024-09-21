package db

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbDriver = os.Getenv("POSTGRES_DRIVER")
	dbUrl    = os.Getenv("POSTGRES_SOURCE")
)

//go:embed migrations/*.sql
var Migrations embed.FS

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbUrl)
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
