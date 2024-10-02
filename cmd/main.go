package main

import (
	v1 "friendly-backend/internal/api/v1/user"
	conn "friendly-backend/internal/db/connection"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/user", v1.GetUserHandler)
	mux.HandleFunc("POST /v1/user", v1.CreateUserHandler)

	db, err := conn.OpenConnection()
	if err != nil {
		panic(err)
	}

	conn.RunMigrations(db)

	http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
