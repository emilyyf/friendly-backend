package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	// TODO:Get values from .env
	dbUser     = os.Getenv("admin")
	dbPassword = os.Getenv("password123")
	dbName     = os.Getenv("friendly")
	dbHost     = os.Getenv("127.0.0.1")
	dbPort     = os.Getenv("6500")
)

func OpenConnection() (*sql.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectionInfo)
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
