package ultrassonography

import (
	"friendly-backend/internal/db/entities/child"
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/medical_history"
	"time"

	"github.com/google/uuid"
)

type Ultrassonography struct {
	ID              uuid.UUID                      `json:"id"`
	Date            time.Time                      `json:"date"`
	Weight          string                         `json:"weight"`
	Height          string                         `json:"height"`
	Percentage      string                         `json:"percentage"`
	BCF             string                         `json:"bcf"`
	ILA             string                         `json:"ila"`
	LIQAM           string                         `json:"liq_am"`
	Placenta        string                         `json:"placenta"`
	Degree          string                         `json:"degree"`
	IDChild         child.Child                    `json:"id_child"`
	IDMedicalHistoy medical_history.MedicalHistory `json:"id_medical_history"`
	UpdateLog       log.Log                        `json:"update_log"`
	CreateLog       log.Log                        `json:"create_log"`
}
