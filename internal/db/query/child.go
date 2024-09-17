package models

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/child"

	"github.com/google/uuid"
)

func InsertChild(child child.Child) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO child (id, id_household, id_mother, name, birth, age, local, race, alive_birth_certificate, alive_birth_certificate,
                            birth_ceriticate, rg, cpf, sus_card, blood_type, weight_at_birth, height_at_birth, first_apgar, fifth_apgar, neonatal_hell_prick,
                            hear_test, hearing_triage, eye_test, od, oe, pregnancy_time, login, msd, mmii, create_log, update_log) VALUES ($1, $2 , $3 , $4 , $5, $6, $7, $8, $9, $10,
                            $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31) RETURN id`

	err = conn.QueryRow(sql, child.ID, child.IDHousehold, child.IDMother, child.Name, child.Birth, child.Age, child.Local, child.Race, child.AliveBirthCertificate, child.BirthCertificate,
		child.RG, child.CPF, child.SUSCard, child.BloodType, child.WeightAtBirth, child.HeightAtBirth, child.FirstApgar, child.FifthApgar, child.NeonatalHeelPrick, child.HearTest,
		child.HearingTriage, child.EyeTest, child.OD, child.OE, child.PregnancyTime, child.Login, child.MSD, child.MMII, child.CreateLog, child.UpdateLog).Scan(&id)

	return
}
