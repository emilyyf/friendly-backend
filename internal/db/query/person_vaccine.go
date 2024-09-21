package db

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/person_vaccine"

	"github.com/google/uuid"
)

func InsertPersonVaccine(person_vaccine person_vaccine.PersonVaccine) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO person_vaccine (id, person_id, vaccine_id, dosage_id, date, create_log, update_log, medical_history_id)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = conn.QueryRow(sql, person_vaccine.ID, person_vaccine.PersonID, person_vaccine.VaccineID, person_vaccine.DosageID,
		person_vaccine.Date, person_vaccine.CreateLog, person_vaccine.UpdateLog, person_vaccine.MedicalHistoryID).Scan(&id)

	return
}

func DeletePersonVaccine(person_vaccine person_vaccine.PersonVaccine) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM person_vaccine WHERE id = $1`

	result, err := conn.Exec(sql, person_vaccine.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected > 0 {
		id = person_vaccine.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetPersonVaccine(person_vaccine person_vaccine.PersonVaccine) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM person_vaccine WHERE ID = $1`, person_vaccine.ID)

	err = row.Scan(&person_vaccine.ID, &person_vaccine.PersonID, &person_vaccine.VaccineID, &person_vaccine.DosageID, &person_vaccine.Date, &person_vaccine.CreateLog, &person_vaccine.UpdateLog, &person_vaccine.MedicalHistoryID)

	return
}

func GetAllPersonVaccine() (person_vaccines []person_vaccine.PersonVaccine, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM person_vaccine`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var person_vaccine person_vaccine.PersonVaccine
		err = rows.Scan(&person_vaccine.ID, &person_vaccine.PersonID, &person_vaccine.VaccineID, &person_vaccine.DosageID, &person_vaccine.Date, &person_vaccine.CreateLog, &person_vaccine.UpdateLog, &person_vaccine.MedicalHistoryID)
		if err != nil {
			return nil, err
		}
		person_vaccines = append(person_vaccines, person_vaccine)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return person_vaccines, nil
}

func UpdatePersonVaccine(person_vaccine person_vaccine.PersonVaccine) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE person_vaccine SET person_id = $2, vaccine_id = $3, dosage_id = $4, date = $5, create_log = $6, update_log = $7, medical_history_id = $8 WHERE id = $1`

	_, err = conn.Exec(sql, person_vaccine.ID, person_vaccine.PersonID, person_vaccine.VaccineID, person_vaccine.DosageID,
		person_vaccine.Date, person_vaccine.CreateLog, person_vaccine.UpdateLog, person_vaccine.MedicalHistoryID)
	if err != nil {
		return uuid.Nil, err
	}
	id = person_vaccine.ID
	return
}
