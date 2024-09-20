package db

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/exams"

	"github.com/google/uuid"
)

func InsertExams(exams exams.Exams) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()
	sql := `INSERT INTO exams (id, description, date, result, igm, igg, id_medical_history, create_log, update_log
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = conn.QueryRow(sql, exams.ID, exams.Description, exams.Date, exams.Result, exams.IGM, exams.IGG, exams.IDMedicalHistory, exams.CreateLog, exams.UpdateLog).Scan(&id)

	return
}

func DeleteExams(exams exams.Exams) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM exams WHERE id = $1`

	result, err := conn.Exec(sql, exams.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = exams.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetExams(exams exams.Exams) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM exams WHERE ID = $1`, exams.ID)

	err = row.Scan(&exams.ID, &exams.Description, &exams.Date, &exams.Result, &exams.IGM, &exams.IGG, &exams.IDMedicalHistory, &exams.CreateLog, &exams.UpdateLog)

	return
}

func GetAllExams() (exam []exams.Exams, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM exams`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var exams exams.Exams
		err = rows.Scan(&exams.ID, &exams.Description, &exams.Date, &exams.Result, &exams.IGM, &exams.IGG, &exams.IDMedicalHistory, &exams.CreateLog, &exams.UpdateLog)
		if err != nil {
			return nil, err
		}
		exam = append(exam, exams)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return exam, nil
}

func UpdateExams(exams exams.Exams) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE exams SET description = $2, date = $3, result = $4, igm = $5, igg = $6, id_medical_history = $7, create_log = $8, update_log = $9`

	_, err = conn.Exec(sql, exams.Description, exams.Date, exams.Result, exams.IGM, exams.IGG, exams.IDMedicalHistory, exams.CreateLog, exams.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = exams.ID
	return
}
