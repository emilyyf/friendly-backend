package vaccines

import (
	"friendly-backend/internal/db/entities/log"

	"github.com/google/uuid"
)

type Vaccines struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreateLog log.Log   `json:"create_log"`
	UpdateLog log.Log   `json:"update_log"`
}
