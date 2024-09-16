package models

import (
	"time"

	"github.com/google/uuid"
)

type PersonVaccine struct {
	ID               uuid.UUID      `json:"id"`
	PersonID         Person         `json:"person_id"`
	VaccineID        Vaccines       `json:"vacinne_id"`
	DosageID         VaccineDosages `json:"dosage_id"`
	Date             time.Time      `json:"date"`
	CreateLog        Log            `json:"create_log"`
	UpdateLog        Log            `json:"update_log"`
	MedicalHistoryID MedicalHistory `json:"medical_history_id"`
}
