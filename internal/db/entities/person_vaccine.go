package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonVaccine struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	PersonID         uuid.UUID `json:"person_id" gorm:"type:uuid;"`
	Person           Person
	VaccineID        uuid.UUID `json:"vacinne_id" gorm:"type:uuid;"`
	Vaccine          Vaccines
	DosageID         uuid.UUID `json:"dosage_id" gorm:"type:uuid;"`
	Dosage           VaccineDosages
	MedicalHistoryID uuid.UUID `json:"medical_history_id" gorm:"type:uuid;"`
	MedicalHistory   MedicalHistory
	Date             time.Time `json:"date"`
	CreateLogID      uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog        Log
	UpdateLogID      uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog        Log
}

func (p *PersonVaccine) BeforeCreate(d *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
