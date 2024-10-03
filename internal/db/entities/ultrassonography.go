package entities

import (
	"time"

	"github.com/google/uuid"
)

type Ultrassonography struct {
	ID              uuid.UUID      `json:"id" gorm:"primaryKey"`
	Date            time.Time      `json:"date"`
	Weight          string         `json:"weight"`
	Height          string         `json:"height"`
	Percentage      string         `json:"percentage"`
	BCF             string         `json:"bcf"`
	ILA             string         `json:"ila"`
	LIQAM           string         `json:"liq_am"`
	Placenta        string         `json:"placenta"`
	Degree          string         `json:"degree"`
	IDChild         Child          `json:"id_child"`
	IDMedicalHistoy MedicalHistory `json:"id_medical_history"`
	UpdateLog       Log            `json:"update_log"`
	CreateLog       Log            `json:"create_log"`
}
