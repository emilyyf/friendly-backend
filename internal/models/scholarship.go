package models

import "github.com/google/uuid"

type Scholarship struct {
	ID      uuid.UUID `json:"id"`
	ChildID Child     `json:"child_id"`
	Grade   string    `json:"grade"`
	School  string    `json:"school"`
	Year    int64     `json:"year"`
	Period  int64     `json:"period"`
}
