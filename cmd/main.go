package main

import (
	v1 "friendly-backend/internal/api/v1/user"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/user", v1.GetUserHandler)
	mux.HandleFunc("POST /v1/user", v1.CreateUserHandler)

	http.ListenAndServe(":8000", mux)
}
