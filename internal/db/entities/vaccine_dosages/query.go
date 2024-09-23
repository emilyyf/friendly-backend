package vaccine_dosages

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertVaccineDosages(vaccine_dosage VaccineDosages) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `INSERT INTO vaccine_dosages (id, vaccine_id, description, create_log, update_log)
                VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = conn.QueryRow(sql, vaccine_dosage.ID, vaccine_dosage.VaccineID, vaccine_dosage.Descripiton, vaccine_dosage.CreateLog, vaccine_dosage.UpdateLog).Scan(&id)

	return
}

func DeletVaccineDosages(vaccine_dosages VaccineDosages) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `DELETE FROM vaccine_dosages WHERE ID = $1`
	result, err := conn.Exec(sql, vaccine_dosages.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected > 0 {
		id = vaccine_dosages.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetVaccineDosages(vaccine_dosages VaccineDosages) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM vaccine_dosages WHERE ID = $1`, vaccine_dosages.ID)

	err = row.Scan(&vaccine_dosages.ID, &vaccine_dosages.VaccineID, &vaccine_dosages.Descripiton, &vaccine_dosages.CreateLog, &vaccine_dosages.UpdateLog)

	return
}

func GetAllVaccineDosages() (vaccine_dosage []VaccineDosages, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT FROM vaccine_dosages`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var vaccine_dosages VaccineDosages
		err = rows.Scan(&vaccine_dosages.ID, &vaccine_dosages.VaccineID, &vaccine_dosages.Descripiton, &vaccine_dosages.CreateLog, &vaccine_dosages.UpdateLog)
		if err != nil {
			return nil, err
		}
		vaccine_dosage = append(vaccine_dosage, vaccine_dosages)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return vaccine_dosage, nil
}

func UpdateVaccineDosages(vaccine_dosages VaccineDosages) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE vaccine_dosages SET vaccine_id = $2, description = $3, create_log = $4, update_log = $5 WHERE id = $1`

	_, err = conn.Exec(sql, vaccine_dosages.ID, vaccine_dosages.VaccineID, vaccine_dosages.Descripiton, vaccine_dosages.CreateLog, vaccine_dosages.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = vaccine_dosages.ID
	return
}
