package models

import "github.com/google/uuid"

type Vaccines struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreateLog Log       `json:"create_log"`
	UpdateLog Log       `json:"update_log"`
}
