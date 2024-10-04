package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ultrassonography struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Date             time.Time `json:"date"`
	Weight           string    `json:"weight"`
	Height           string    `json:"height"`
	Percentage       string    `json:"percentage"`
	BCF              string    `json:"bcf"`
	ILA              string    `json:"ila"`
	LIQAM            string    `json:"liq_am"`
	Placenta         string    `json:"placenta"`
	Degree           string    `json:"degree"`
	ChildID          uuid.UUID `json:"child_id" gorm:"type:uuid;"`
	Child            Child
	MedicalHistoryID uuid.UUID `json:"medical_history_id" gorm:"type:uuid;"`
	MedicalHistory   MedicalHistory
	CreateLogID      uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog        Log
	UpdateLogID      uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog        Log
}

func (u *Ultrassonography) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
