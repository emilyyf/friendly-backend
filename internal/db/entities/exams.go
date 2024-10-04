package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exams struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Description      string    `json:"description"`
	Date             time.Time `json:"date"`
	Result           string    `json:"result"`
	IGM              string    `json:"igm"`
	IGG              string    `json:"igg"`
	MedicalHistoryID uuid.UUID `json:"medical_history_id" gorm:"type:uuid;"`
	MedicalHistory   MedicalHistory
	CreateLogID      uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog        Log
	UpdateLogID      uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog        Log
}

func (e *Exams) BeforeCreate(d *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
