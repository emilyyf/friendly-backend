package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VaccineDosages struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	VaccineID   uuid.UUID `json:"vaccine_id" gorm:"type:uuid;"`
	Vaccine     Vaccines
	Descripiton string    `json:"description"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (v *VaccineDosages) BeforeCreate(d *gorm.DB) (err error) {
	v.ID = uuid.New()
	return
}
