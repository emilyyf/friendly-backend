package db

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
                            $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31) RETURNING id`

	err = conn.QueryRow(sql, child.ID, child.IDHousehold, child.IDMother, child.Name, child.Birth, child.Age, child.Local, child.Race, child.AliveBirthCertificate, child.BirthCertificate,
		child.RG, child.CPF, child.SUSCard, child.BloodType, child.WeightAtBirth, child.HeightAtBirth, child.FirstApgar, child.FifthApgar, child.NeonatalHeelPrick, child.HearTest,
		child.HearingTriage, child.EyeTest, child.OD, child.OE, child.PregnancyTime, child.Login, child.MSD, child.MMII, child.CreateLog, child.UpdateLog).Scan(&id)

	return
}

func DeleteChild(child child.Child) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `DELETE FROM child where id = $1`

	result, err := conn.Exec(sql, child.ID)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected > 0 {
		id = child.ID
	} else {
		id = uuid.Nil
	}

	return
}

func GetChild(child child.Child) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM child WHERE id=$1`, child.ID)

	err = row.Scan(&child.ID, &child.IDHousehold, &child.IDMother, &child.Name, &child.Birth, &child.Age, &child.Local, &child.Race, &child.AliveBirthCertificate, &child.BirthCertificate,
		&child.RG, &child.CPF, &child.SUSCard, &child.BloodType, &child.WeightAtBirth, &child.HeightAtBirth, &child.FirstApgar, &child.FifthApgar, &child.NeonatalHeelPrick, &child.HearTest,
		&child.HearingTriage, &child.EyeTest, &child.OD, &child.OE, &child.PregnancyTime, &child.Login, &child.MSD, &child.MMII, &child.CreateLog, &child.UpdateLog)

	return
}

func GetAllChild() (children []child.Child, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM child`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var child child.Child
		err = rows.Scan(&child.ID, &child.IDHousehold, &child.IDMother, &child.Name, &child.Birth, &child.Age, &child.Local, &child.Race, &child.AliveBirthCertificate, &child.BirthCertificate,
			&child.RG, &child.CPF, &child.SUSCard, &child.BloodType, &child.WeightAtBirth, &child.HeightAtBirth, &child.FirstApgar, &child.FifthApgar, &child.NeonatalHeelPrick, &child.HearTest,
			&child.HearingTriage, &child.EyeTest, &child.OD, &child.OE, &child.PregnancyTime, &child.Login, &child.MSD, &child.MMII, &child.CreateLog, &child.UpdateLog)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return children, nil
}

func UpdateChild(child child.Child) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `UPDATE child SET id_household = $2 , id_mother = $3 , name =  $4 , birth =  $5, age = $6, local = $7, race =  $8, alive_birth_certificate =  $9, birth_ceriticate = $10,
                            rg = $11, cpf =  $12, sus_card = $13, blood_type = $14, weight_at_birth = $15, height_at_birth = $16, first_apgar = $17, fifth_apgar = $18, neonatal_hell_prick = $19, hear_test = $20, hearing_triage = $21,
                            eye_test = $22, od = $23, oe = $24, pregnancy_time = $25, login = $26, msd = $27, mmii = $28, create_log = $29, update_log = $30 WHERE id = $1`

	_, err = conn.Exec(sql, child.ID, child.IDHousehold, child.IDMother, child.Name, child.Birth, child.Age, child.Local, child.Race, child.AliveBirthCertificate, child.BirthCertificate,
		child.RG, child.CPF, child.SUSCard, child.BloodType, child.WeightAtBirth, child.HeightAtBirth, child.FirstApgar, child.FifthApgar, child.NeonatalHeelPrick, child.HearTest, child.HearingTriage,
		child.EyeTest, child.OD, child.OE, child.PregnancyTime, child.Login, child.MSD, child.MMII, child.CreateLog, child.UpdateLog)
	if err != nil {
		return
	}

	id = child.ID
	return
}
