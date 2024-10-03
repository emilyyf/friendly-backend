package entities

import (
	"github.com/google/uuid"
)

type VaccineDosages struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	VaccineID   Vaccines  `json:"vaccine_id"`
	Descripiton string    `json:"description"`
	CreateLog   Log       `json:"create_log"`
	UpdateLog   Log       `json:"update_log"`
}
