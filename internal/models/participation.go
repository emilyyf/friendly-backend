package models

import (
	"time"

	"github.com/google/uuid"
)

type Participation struct {
	ID          uuid.UUID `json:"id"`
	IDChild     Child     `json:"id_child"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreateLog   Log       `json:"create_log"`
	UpdateLog   Log       `json:"update_log"`
}
