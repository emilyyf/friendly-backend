package person

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertPerson(person Person) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO person (id, id_household, birth, age, name, country, uf, phone, graduation, rg, rg_exp, cpf, 
                                    sus_card, card_series, card_uf, company, work_function, enployment_card, hiring_date,
                                    resignation_date, salary, extra_income, create_log, update_log)
                                    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
                                    $17, $18, $19, $20, $21, $22, $23, $24, $25) RETURNING id`

	err = conn.QueryRow(sql, person.ID, person.IDHousehold, person.Birth, person.Age, person.Name, person.Country, person.UF, person.Phone,
		person.Graduation, person.RG, person.RGEXP, person.CPF, person.CardSeries, person.CardUF, person.Company, person.WorkFunction,
		person.EnploymentCard, person.HiringDate, person.ResignationDate, person.Salary, person.ExtraIncome, person.CreateLog, person.UpdateLog).Scan(&id)

	return
}

func DeletePerson(person Person) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM person WHERE id = $1`

	result, err := conn.Exec(sql, person.ID)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected > 0 {
		id = person.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetPerson(person Person) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM person WHERE ID = $1`, person.ID)

	err = row.Scan(&person.ID, &person.IDHousehold, &person.Birth, &person.Age, &person.Name, &person.Country, &person.UF, &person.Phone,
		&person.Graduation, &person.RG, &person.RGEXP, &person.CPF, &person.CardSeries, &person.CardUF, &person.Company, &person.WorkFunction,
		&person.EnploymentCard, &person.HiringDate, &person.ResignationDate, &person.Salary, &person.ExtraIncome, &person.CreateLog, &person.UpdateLog)

	return
}

func GetAllPerson() (persons []Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM person`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var person Person
		err = rows.Scan(&person.ID, &person.IDHousehold, &person.Birth, &person.Age, &person.Name, &person.Country, &person.UF, &person.Phone,
			&person.Graduation, &person.RG, &person.RGEXP, &person.CPF, &person.CardSeries, &person.CardUF, &person.Company, &person.WorkFunction,
			&person.EnploymentCard, &person.HiringDate, &person.ResignationDate, &person.Salary, &person.ExtraIncome, &person.CreateLog, &person.UpdateLog)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return persons, nil
}

func UpdatePersonVaccine(person Person) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	sql := `UPDATE person SET id_household = $2, birth = $3, age = $4, name = $5, country = $6, uf = $7, phone = $8, graduation = $9, rg = $10, rg_exp = $11, cpf = $12, 
                                    sus_card = $13, card_series = $14, card_uf = $15, company = $16, work_function = $17, enployment_card = $18, hiring_date = $19,
                                    resignation_date = $20, salary = $21, extra_income = $22, create_log = $23, update_log = $24 WHERE id = $1`

	_, err = conn.Exec(sql, person.ID, person.IDHousehold, person.Birth, person.Age, person.Name, person.Country, person.UF, person.Phone,
		person.Graduation, person.RG, person.RGEXP, person.CPF, person.CardSeries, person.CardUF, person.Company, person.WorkFunction,
		person.EnploymentCard, person.HiringDate, person.ResignationDate, person.Salary, person.ExtraIncome, person.CreateLog, person.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}
	id = person.ID
	return
}
