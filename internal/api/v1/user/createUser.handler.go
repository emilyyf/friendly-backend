package v1

import (
	"encoding/json"
	"friendly-backend/internal/models"
	"io"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var data models.UserResponse
	var ret []byte

	body, err := io.ReadAll(r.Body)
	if err != nil {
		goto error
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		goto error
	}

	data.Name = "Placeholder"

	ret, err = json.Marshal(data)
	if err != nil {
		goto error
	}

	w.Write(ret)

	return

error:
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Invalid request!"))
}
