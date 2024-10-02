package person_vaccine

import (
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/medical_history"
	"friendly-backend/internal/db/entities/person"
	"friendly-backend/internal/db/entities/vaccine_dosages"
	"friendly-backend/internal/db/entities/vaccines"
	"time"

	"github.com/google/uuid"
)

type PersonVaccine struct {
	ID               uuid.UUID                      `json:"id" gorm:"primaryKey"`
	PersonID         person.Person                  `json:"person_id"`
	VaccineID        vaccines.Vaccines              `json:"vacinne_id"`
	DosageID         vaccine_dosages.VaccineDosages `json:"dosage_id"`
	Date             time.Time                      `json:"date"`
	CreateLog        log.Log                        `json:"create_log"`
	UpdateLog        log.Log                        `json:"update_log"`
	MedicalHistoryID medical_history.MedicalHistory `json:"medical_history_id"`
}
