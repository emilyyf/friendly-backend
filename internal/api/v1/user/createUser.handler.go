package v1

import "net/http"

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
