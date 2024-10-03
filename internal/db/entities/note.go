package entities

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	IDPerson    Person    `json:"id_person"`
	IDChild     Child     `json:"id_child"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreateLog   Log       `json:"create_log"`
	UpdateLog   Log       `json:"update_log"`
}
