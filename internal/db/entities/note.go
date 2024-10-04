package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	PersonID    uuid.UUID `json:"person_id" gorm:"type:uuid;"`
	Person      Person
	ChildID     uuid.UUID `json:"child_id" gorm:"type:uuid;"`
	Child       Child
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreateLogID uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog   Log
	UpdateLogID uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog   Log
}

func (n *Note) BeforeCreate(d *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}
