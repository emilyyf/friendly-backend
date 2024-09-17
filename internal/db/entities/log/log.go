package log

import (
	"friendly-backend/internal/models"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID     uuid.UUID           `json:"id"`
	UserID models.UserResponse `json:"user_id"`
	Table  string              `json:"table"`
	Date   time.Time           `json:"date"`
}
