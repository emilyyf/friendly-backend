package models

import (
	"time"

	"github.com/google/uuid"
)

type Exams struct {
	ID               uuid.UUID      `json:"id"`
	Description      string         `json:"description"`
	Date             time.Time      `json:"date"`
	Result           string         `json:"result"`
	IGM              string         `json:"igm"`
	IGG              string         `json:"igg"`
	IDMedicalHistory MedicalHistory `json:"id_medical_history"`
	CreateLog        Log            `json:"create_log"`
	UpdateLog        Log            `json:"update_log"`
}
