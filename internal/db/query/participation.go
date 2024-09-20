package db

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/participation"

	"github.com/google/uuid"
)

func InsertParticipation(participation participation.Participation) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERTO INTO participation (id, id_child, date, description, create_log, update_log
                  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(sql, participation.ID, participation.IDChild, participation.Date, participation.Description, participation.CreateLog, participation.UpdateLog).Scan(&id)

	return
}

func DeleteParticipation(participation participation.Participation) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM participation WHERE id = $1`

	result, err := conn.Exec(sql, participation.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = participation.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetParticipation(participation participation.Participation) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM participation WHERE ID = $1`, participation.ID)

	err = row.Scan(&participation.ID, &participation.IDChild, &participation.Date, &participation.Description, &participation.CreateLog, &participation.UpdateLog)

	return
}

func GetAllParticipation() (participations []participation.Participation, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM participation`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var participation participation.Participation
		err = rows.Scan(&participation.ID, &participation.IDChild, &participation.Date, &participation.Description, &participation.CreateLog, &participation.UpdateLog)
		if err != nil {
			return nil, err
		}
		participations = append(participations, participation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return participations, nil
}

func UpdateParticipations(participation participation.Participation) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `UPDATE participation SET 
                id_child = $2,
                date = $3, 
                description = $4,
                create_log = $5,
                update_log = $6
                WHERE id = $1`

	_, err = conn.Exec(sql, participation.ID, participation.IDChild, participation.Date, participation.Description, participation.CreateLog, participation.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}

	id = participation.ID
	return
}
