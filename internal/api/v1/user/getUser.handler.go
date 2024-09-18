package v1

import (
	"encoding/json"
	"friendly-backend/internal/models"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, h *http.Request) {
	var data models.UserResponse
	data.Name = "Placeholder"

	ret, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request!"))
	}

	w.Write(ret)
}
