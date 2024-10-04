package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participation struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	ChildID     uuid.UUID `json:"child_id" gorm:"type:uuid;"`
	Child       Child
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (p *Participation) BeforeCreate(d *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
