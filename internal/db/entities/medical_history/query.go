package medical_history

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertMedicalHistory(medical_history MedicalHistory) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `INSERT INTO medical_history (id, id_person, smoker, alcohool, aborts, vaginal_delivery, caesarian, pregnancy, blood_type, blood_glucose, syphilis, hiv, create_log, update_log) 
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning ID`

	err = conn.QueryRow(sql, medical_history.ID, medical_history.IDPerson, medical_history.Smoker, medical_history.Alcohool, medical_history.Aborts, medical_history.VaginalDeleviry,
		medical_history.Caesarian, medical_history.Pregnancy, medical_history.BloodType, medical_history.BloodGlucose, medical_history.Syphilis, medical_history.HIV, medical_history.CreateLog, medical_history.UpdateLog).Scan(&id)

	return
}

func DeleteMedicalHistory(medical_history MedicalHistory) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `DELETE FROM medical_history WHERE ID = $1`

	result, err := conn.Exec(sql, medical_history.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected > 0 {
		id = medical_history.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetMedicalHistory(medical_history MedicalHistory) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM medical_history WHERE ID = $1`, medical_history.ID)

	err = row.Scan(&medical_history.ID, &medical_history.IDPerson, &medical_history.Smoker, &medical_history.Alcohool, &medical_history.Aborts, &medical_history.VaginalDeleviry,
		&medical_history.Caesarian, &medical_history.Pregnancy, &medical_history.BloodType, &medical_history.BloodGlucose, &medical_history.Syphilis, &medical_history.HIV, &medical_history.CreateLog, &medical_history.UpdateLog)

	return
}

func GetAllMedicalHistory() (medical_historys []MedicalHistory, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM medical_history`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var medical_history MedicalHistory
		err = rows.Scan(&medical_history.ID, &medical_history.IDPerson, &medical_history.Smoker, &medical_history.Alcohool, &medical_history.Aborts, &medical_history.VaginalDeleviry,
			&medical_history.Caesarian, &medical_history.Pregnancy, &medical_history.BloodType, &medical_history.BloodGlucose, &medical_history.Syphilis, &medical_history.HIV, &medical_history.CreateLog, &medical_history.UpdateLog)
		if err != nil {
			return nil, err
		}
		medical_historys = append(medical_historys, medical_history)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return medical_historys, nil
}

func UpdateMedicalHistory(medical_history MedicalHistory) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE medical_history SET id_person = $2, smoker = $3, $alcohool = $4, aborts = $5, vaginal_delivery = $6, caesarian = $7, pregnancy = $8, 
                                        blood_type = $9, blood_glucose = $10, syphilis = $11, hiv = $12, create_log = $13, update_log = $14 WHERE id = $1`

	_, err = conn.Exec(sql, medical_history.ID, medical_history.IDPerson, medical_history.Smoker, medical_history.Alcohool, medical_history.Aborts, medical_history.VaginalDeleviry,
		medical_history.Caesarian, medical_history.Pregnancy, medical_history.BloodType, medical_history.BloodGlucose, medical_history.Syphilis, medical_history.HIV, medical_history.CreateLog, medical_history.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = medical_history.ID
	return
}
