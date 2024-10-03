package entities

import (
	"github.com/google/uuid"
)

type Scholarship struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	ChildID Child     `json:"child_id"`
	Grade   string    `json:"grade"`
	School  string    `json:"school"`
	Year    int64     `json:"year"`
	Period  int64     `json:"period"`
}
