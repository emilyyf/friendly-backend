package models

import (
	"friendly-backend/internal/models"

	"github.com/google/uuid"
)

func InsertChild(child models.Child) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO child (id, id_household, id_mother, name, birth, age, local, race, alive_birth_certificate, alive_birth_certificate,
                            birth_ceriticate, rg, cpf, sus_card, blood_type, weight_at_birth, height_at_birth, first_apgar, fifth_apgar, neonatal_hell_prick,
                            hear_test, hearing_triage, eye_test, od, oe, pregnancy_time, login, msd, mmii, create_log, update_log)`

	return
}
