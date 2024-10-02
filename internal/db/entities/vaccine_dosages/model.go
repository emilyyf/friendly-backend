package vaccine_dosages

import (
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/vaccines"

	"github.com/google/uuid"
)

type VaccineDosages struct {
	ID          uuid.UUID         `json:"id" gorm:"primaryKey"`
	VaccineID   vaccines.Vaccines `json:"vaccine_id"`
	Descripiton string            `json:"description"`
	CreateLog   log.Log           `json:"create_log"`
	UpdateLog   log.Log           `json:"update_log"`
}
