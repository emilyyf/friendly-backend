package entities

import (
	"github.com/google/uuid"
)

type MedicalHistory struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	IDPerson        Person    `json:"id_person"`
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
	CreateLog       Log       `json:"create_log"`
	UpdateLog       Log       `json:"update_log"`
}
