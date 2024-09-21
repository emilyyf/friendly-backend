package appointment

import (
	"friendly-backend/internal/db/entities/child"
	"friendly-backend/internal/db/entities/log"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID   `json:"id"`
	IDChild   child.Child `json:"id_child"`
	Date      time.Time   `json:"date"`
	IG        string      `json:"ig"`
	Weight    string      `json:"weight"`
	PA        string      `json:"pa"`
	AU        string      `json:"au"`
	BCF       string      `json:"bcf"`
	CreateLog log.Log     `json:"create_log"`
	UpdateLog log.Log     `json:"update_log"`
}
