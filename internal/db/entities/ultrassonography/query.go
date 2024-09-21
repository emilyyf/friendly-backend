package ultrassonography

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertUltrassonography(ultrassonography Ultrassonography) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO ultrassonography (id, date, weight, height, percentage, bcf, ila, liq_am, placenta, 
                                              degree, id_child, id_medical_history, update_log, create_log)
                                              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	err = conn.QueryRow(sql, ultrassonography.ID, ultrassonography.Date, ultrassonography.Weight, ultrassonography.Height,
		ultrassonography.Percentage, ultrassonography.BCF, ultrassonography.ILA, ultrassonography.LIQAM,
		ultrassonography.Placenta, ultrassonography.Degree, ultrassonography.IDChild, ultrassonography.IDMedicalHistoy,
		ultrassonography.UpdateLog, ultrassonography.CreateLog).Scan(&id)

	return
}

func DeleteUltrassonography(ultrassonography Ultrassonography) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM ultrassonography WHERE ID = $1`

	result, err := conn.Exec(sql, ultrassonography.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = ultrassonography.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetUltrassonography(ultrassonography Ultrassonography) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM ultrassonography WHERE id = $1`, ultrassonography.ID)

	err = row.Scan(&ultrassonography.ID, &ultrassonography.Date, &ultrassonography.Weight, &ultrassonography.Height,
		&ultrassonography.Percentage, &ultrassonography.BCF, &ultrassonography.ILA, &ultrassonography.LIQAM,
		&ultrassonography.Placenta, &ultrassonography.Degree, &ultrassonography.IDChild, &ultrassonography.IDMedicalHistoy,
		&ultrassonography.UpdateLog, &ultrassonography.CreateLog)

	return
}

func GetAllUltrassonography() (ultrassonographys []Ultrassonography, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM ultrassonography`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ultrassonography Ultrassonography
		err = rows.Scan(&ultrassonography.ID, &ultrassonography.Date, &ultrassonography.Weight, &ultrassonography.Height,
			&ultrassonography.Percentage, &ultrassonography.BCF, &ultrassonography.ILA, &ultrassonography.LIQAM,
			&ultrassonography.Placenta, &ultrassonography.Degree, &ultrassonography.IDChild, &ultrassonography.IDMedicalHistoy,
			&ultrassonography.UpdateLog, &ultrassonography.CreateLog)
		if err != nil {
			return nil, err
		}
		ultrassonographys = append(ultrassonographys, ultrassonography)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ultrassonographys, nil
}

func UpdateUltrassonography(ultrassonography Ultrassonography) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `UPDATE ultrassonography SET date = $2, weight = $3, height = $4, percentage = $5, bcf = $6, ila = $7, liq_am = $8,
                                              placenta = $9, degree = $10, id_child = $11, id_medical_history = $12, update_log = $13
                                              create_log = $14 WHERE id = $1`

	_, err = conn.Exec(sql, ultrassonography.ID, ultrassonography.Date, ultrassonography.Weight, ultrassonography.Height,
		ultrassonography.Percentage, ultrassonography.BCF, ultrassonography.ILA, ultrassonography.LIQAM,
		ultrassonography.Placenta, ultrassonography.Degree, ultrassonography.IDChild, ultrassonography.IDMedicalHistoy,
		ultrassonography.UpdateLog, ultrassonography.CreateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = ultrassonography.ID
	return
}
