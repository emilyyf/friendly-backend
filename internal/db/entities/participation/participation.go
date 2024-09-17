package participation

import (
	"friendly-backend/internal/db/entities/child"
	"friendly-backend/internal/db/entities/log"
	"time"

	"github.com/google/uuid"
)

type Participation struct {
	ID          uuid.UUID   `json:"id"`
	IDChild     child.Child `json:"id_child"`
	Date        time.Time   `json:"date"`
	Description string      `json:"description"`
	CreateLog   log.Log     `json:"create_log"`
	UpdateLog   log.Log     `json:"update_log"`
}
