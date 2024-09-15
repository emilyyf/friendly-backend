package v1

import "net/http"

func GetUserHandler(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("Hello"))
}
