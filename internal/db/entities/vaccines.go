package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vaccines struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Name        string    `json:"name"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (v *Vaccines) BeforeCreate(d *gorm.DB) (err error) {
	v.ID = uuid.New()
	return
}
