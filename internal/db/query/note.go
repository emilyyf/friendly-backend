package db

import (
	db "friendly-backend/internal/db/connection"
	"friendly-backend/internal/db/entities/note"

	"github.com/google/uuid"
)

func InsertNote(note note.Note) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO note (id, id_person, id_child, date, description, create_log, update_log) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = conn.QueryRow(sql, note.ID, note.IDPerson, note.IDChild, note.Date, note.Description, note.CreateLog, note.UpdateLog).Scan(&id)

	return
}

func DeleteNote(note note.Note) (id uuid.UUID, err errpr) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM note WHERE id = $1`

	result, err := conn.Exec(sql, note.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = note.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetNote(note note.Note) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM note WHERE ID=$1`, note.ID)

	err = row.Scan(&note.ID, &note.IDPerson, &note.IDChild, &note.Date, &note.Description, &note.CreateLog, &note.UpdateLog)

	return
}

func GetAllNote() (notes []note.Note, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM note`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note note.Note
		err = rows.Scan(&note.ID, &note.IDPerson, &note.IDChild, &note.Date, &note.Description, &note.CreateLog, &note.UpdateLog)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}

func UpdateNote(note note.Note) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE note SET id_person = $2, id_child = $3, date = $ 4, description = $6, create_log = $7, update_log = $8 WHERE id = $1`

	_, err = conn.Exec(sql, note.ID, note.IDPerson, note.IDChild, note.Date, note.Description, note.CreateLog, note.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}

	id = note.ID
	return
}
