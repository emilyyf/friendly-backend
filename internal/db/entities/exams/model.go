package exams

import (
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/medical_history"
	"time"

	"github.com/google/uuid"
)

type Exams struct {
	ID               uuid.UUID                      `json:"id"`
	Description      string                         `json:"description"`
	Date             time.Time                      `json:"date"`
	Result           string                         `json:"result"`
	IGM              string                         `json:"igm"`
	IGG              string                         `json:"igg"`
	IDMedicalHistory medical_history.MedicalHistory `json:"id_medical_history"`
	CreateLog        log.Log                        `json:"create_log"`
	UpdateLog        log.Log                        `json:"update_log"`
}
