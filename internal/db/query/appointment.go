package db

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/appointment"

	"github.com/google/uuid"
)

func InsertAppointment(appointment appointment.Appointment) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO appointment (id, id_child, date, ig, weight, pa, au, bcf, create_log, update_log
                  VALUES ($1. $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err = conn.QueryRow(sql, appointment.ID, appointment.IDChild, appointment.Date, appointment.IG, appointment.Weight, appointment.PA, appointment.AU, appointment.BCF, appointment.CreateLog, appointment.UpdateLog).Scan(&id)

	return
}

func DeleteAppointment(appointment appointment.Appointment) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM appointment WHERE id = $1`

	result, err := conn.Exec(sql, appointment.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected > 0 {
		id = appointment.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetAppointment(appointment appointment.Appointment) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM appointment WHERE ID = $1`, appointment.ID)

	err = row.Scan(&appointment.ID, &appointment.IDChild, &appointment.Date, &appointment.IG, &appointment.Weight, &appointment.PA, &appointment.AU, &appointment.BCF, &appointment.CreateLog, &appointment.UpdateLog)

	return
}

func GetAllApointment() (appointments []appointment.Appointment, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM appointment`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var appointment appointment.Appointment
		err = rows.Scan(&appointment.ID, &appointment.IDChild, &appointment.Date, &appointment.IG, &appointment.Weight, &appointment.PA, &appointment.AU, &appointment.BCF, &appointment.CreateLog, &appointment.UpdateLog)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}

func UpdateAppointment(appointment appointment.Appointment) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE appointment SET id_child = $2, date = $3, ig = $4, weight = $5, pa = $6, au = $7, bcf = $8, create_log = $9, update_log = $10`

	_, err = conn.Exec(sql, appointment.IDChild, appointment.Date, appointment.IG, appointment.Weight, appointment.PA, appointment.BCF, appointment.CreateLog, appointment.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = appointment.ID
	return
}
