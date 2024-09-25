package vaccines

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertVaccines(vaccines Vaccines) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO vaccines (id, name, create_log, update_log) VALUES
                                VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, vaccines.ID, vaccines.Name, vaccines.CreateLog, vaccines.UpdateLog).Scan(&id)

	return
}

func DeleteVaccines(vaccines Vaccines) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `DELETE FROM vaccines WHERE id = $1`

	result, err := conn.Exec(sql, vaccines.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffect > 0 {
		id = vaccines.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetVaccines(vaccines Vaccines) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM vaccines WHERE id = $1`, vaccines.ID)

	err = row.Scan(&vaccines.ID, vaccines.Name, vaccines.CreateLog, vaccines.UpdateLog)

	return
}

func GetAllVaccines() (vaccine []Vaccines, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM vaccines`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vaccines Vaccines
		err = rows.Scan(&vaccines.ID, &vaccines.Name, &vaccines.CreateLog, &vaccines.UpdateLog)
		if err != nil {
			return nil, err
		}
		vaccine = append(vaccine, vaccines)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return vaccine, nil
}

func UpdateVaccines(vaccines Vaccines) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE vaccines SET name = $2, create_log = $3, update_log = $4 WHERE id = $1`

	_, err = conn.Exec(sql, vaccines.ID, vaccines.Name, vaccines.CreateLog, vaccines.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = vaccines.ID
	return
}
