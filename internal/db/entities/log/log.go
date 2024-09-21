package log

import (
	"friendly-backend/internal/db/entities/user"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID         `json:"id"`
	UserID      user.UserResponse `json:"user_id"`
	Table       string            `json:"table"`
	Date        time.Time         `json:"date"`
	Description string            `json:"description"`
	Action      string            `json:"action"`
	RowID       uuid.UUID         `json:"row_id"`
}
