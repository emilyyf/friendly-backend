package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalHistory struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	PersonID        Person    `json:"person_id" gorm:"type:uuid;"`
	Person          Person
	Smoker          bool      `json:"smoker"`
	Alcohool        bool      `json:"alcohool"`
	Aborts          int64     `json:"aborts"`
	VaginalDeleviry int64     `json:"vaginal_delivery"`
	Caesarian       int64     `json:"caesarian"`
	Pregnancy       string    `json:"pregnancy"`
	BloodType       string    `json:"blood_type"`
	BloodGlucose    string    `json:"blood_glucose"`
	Syphilis        string    `json:"syphilis"`
	HIV             string    `json:"hiv"`
	CreateLogID     uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog       Log
	UpdateLogID     uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog       Log
}

func (m *MedicalHistory) BeforeCreate(d *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
