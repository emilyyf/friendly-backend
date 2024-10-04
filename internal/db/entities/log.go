package entities

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	User        User
	Table       string    `json:"table"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Action      string    `json:"action"`
	RowID       uuid.UUID `json:"row_id"`
}

func (l *Log) BeforeCreate(d *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
