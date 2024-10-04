package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	ChildID     uuid.UUID `json:"child_id" gorm:"type:uuid;"`
	Child       Child
	Date        time.Time `json:"date"`
	IG          string    `json:"ig"`
	Weight      string    `json:"weight"`
	PA          string    `json:"pa"`
	AU          string    `json:"au"`
	BCF         string    `json:"bcf"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (a *Appointment) BeforeCreate(d *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
