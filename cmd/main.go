package main

import (
	v1 "friendly-backend/internal/api/v1/user"
	conn "friendly-backend/internal/db/connection"
	"net/http"

	"github.com/pressly/goose/v3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/user", v1.GetUserHandler)
	mux.HandleFunc("POST /v1/user", v1.CreateUserHandler)

	db, err := conn.OpenConnection()
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(conn.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", mux)
}
