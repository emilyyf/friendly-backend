package note

import (
	"friendly-backend/internal/db/entities/child"
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/person"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID          uuid.UUID     `json:"id"`
	IDPerson    person.Person `json:"id_person"`
	IDChild     child.Child   `json:"id_child"`
	Date        time.Time     `json:"date"`
	Description string        `json:"description"`
	CreateLog   log.Log       `json:"create_log"`
	UpdateLog   log.Log       `json:"update_log"`
}
