package scholarship

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertScholarship(scholarship Scholarship) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO scholarship (id, child_id, grade, school, year, period) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(sql, scholarship.ID, scholarship.ChildID, scholarship.Grade, scholarship.School, scholarship.Year, scholarship.Period).Scan(&id)

	return
}

func DeleteScholarship(scholarship Scholarship) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM scholarship WHERE id = $1`

	result, err := conn.Exec(sql, scholarship.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = scholarship.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetScholarship(scholarship Scholarship) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM scholarship WHERE id = $1`, scholarship.ID)

	err = row.Scan(&scholarship.ID, &scholarship.ChildID, &scholarship.Grade, &scholarship.School, &scholarship.Year, &scholarship.Period)

	return
}

func GetAllScholarship() (scholarships []Scholarship, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM scholarship`)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	for rows.Next() {
		var scholarship Scholarship
		err = rows.Scan(&scholarship.ID, &scholarship.ChildID, &scholarship.Grade, &scholarship.School, &scholarship.Year, &scholarship.Period)
		if err != nil {
			return nil, err
		}
		scholarships = append(scholarships, scholarship)
	}
	if err != nil {
		return nil, err
	}
	return scholarships, nil
}

func UpdateScholarship(scholarship Scholarship) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE scholarship SET
          child_id = $2,
          grade = $3,
          school = $4,
          year = $5,
          period = $6`

	_, err = conn.Exec(sql, scholarship.ID, scholarship.ChildID, scholarship.Grade, scholarship.School, scholarship.Year, scholarship.Period)
	if err != nil {
		return uuid.Nil, err
	}
	id = scholarship.ID
	return
}
