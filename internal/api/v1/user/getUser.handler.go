package v1

import (
	"encoding/json"
	"friendly-backend/internal/db/entities/user"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, h *http.Request) {
	var data user.UserResponse
	data.Name = "Placeholder"

	ret, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request!"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
}
