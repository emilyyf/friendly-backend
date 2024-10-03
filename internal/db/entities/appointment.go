package entities

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	IDChild   Child     `json:"id_child"`
	Date      time.Time `json:"date"`
	IG        string    `json:"ig"`
	Weight    string    `json:"weight"`
	PA        string    `json:"pa"`
	AU        string    `json:"au"`
	BCF       string    `json:"bcf"`
	CreateLog Log       `json:"create_log"`
	UpdateLog Log       `json:"update_log"`
}
