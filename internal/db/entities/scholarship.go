package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Scholarship struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	ChildID     uuid.UUID `json:"child_id" gorm:"type:uuid;"`
	Child       Child
	Grade       string    `json:"grade"`
	School      string    `json:"school"`
	Year        int64     `json:"year"`
	Period      int64     `json:"period"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (s *Scholarship) BeforeCreate(d *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
