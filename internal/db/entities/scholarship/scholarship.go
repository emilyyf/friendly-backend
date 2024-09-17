package scholarship

import (
	"friendly-backend/internal/db/entities/child"

	"github.com/google/uuid"
)

type Scholarship struct {
	ID      uuid.UUID   `json:"id"`
	ChildID child.Child `json:"child_id"`
	Grade   string      `json:"grade"`
	School  string      `json:"school"`
	Year    int64       `json:"year"`
	Period  int64       `json:"period"`
}
