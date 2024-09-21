package household

import (
	db "friendly-backend/internal/db/connection"

	"github.com/google/uuid"
)

func InsertHousehold(household Household) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `INSERT INTO household (id, code, date, adress, adress_number, adress_complement, cep, city, 
                neighborhood, residense_type, rent_value, building_materials, in_house_bathroom, residents,
                rooms, beds, car, television, refrigerator, microwave, road_type, refering_person, create_log, update_log)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $18, $19, $20,
                $21, $22, $23, $24, $25) RETURNING id`

	err = conn.QueryRow(sql, household.ID, household.Code, household.Date, household.Adress, household.AdressNumber,
		household.AdressComplement, household.CEP, household.City, household.Neighborhood, household.ResidenseType,
		household.RentValue, household.BuildingMaterials, household.InHouseBathroom, household.Residents, household.Rooms,
		household.Beds, household.Car, household.Television, household.Refrigerator, household.Microwave, household.RoadType,
		household.ReferingPerson, household.CreateLog, household.UpdateLog).Scan(&id)

	return
}

func DeleteHousehold(household Household) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `DELETE FROM household WHERE id = $1`

	result, err := conn.Exec(sql, household.ID)
	if err != nil {
		return uuid.Nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected > 0 {
		id = household.ID
	} else {
		id = uuid.Nil
	}
	return
}

func GetHousehold(household Household) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM household WHERE id = $1`, household.ID)

	err = row.Scan(&household.ID, &household.Code, &household.Date, &household.Adress, &household.AdressNumber,
		&household.AdressComplement, &household.CEP, &household.City, &household.Neighborhood, &household.ResidenseType,
		&household.RentValue, &household.BuildingMaterials, &household.InHouseBathroom, &household.Residents, &household.Rooms,
		&household.Beds, &household.Car, &household.Television, &household.Refrigerator, &household.Microwave, &household.RoadType,
		&household.ReferingPerson, &household.CreateLog, &household.UpdateLog)

	return
}

func GetAllHousehold() (households []Household, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM household`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var household Household
		err = rows.Scan(&household.ID, &household.Code, &household.Date, &household.Adress, &household.AdressNumber,
			&household.AdressComplement, &household.CEP, &household.City, &household.Neighborhood, &household.ResidenseType,
			&household.RentValue, &household.BuildingMaterials, &household.InHouseBathroom, &household.Residents, &household.Rooms,
			&household.Beds, &household.Car, &household.Television, &household.Refrigerator, &household.Microwave, &household.RoadType,
			&household.ReferingPerson, &household.CreateLog, &household.UpdateLog)
		if err != nil {
			return nil, err
		}
		households = append(households, household)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return households, nil
}

func UpdateHousehold(household Household) (id uuid.UUID, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return uuid.Nil, err
	}

	defer conn.Close()

	sql := `UPDATE household SET code = $2, date = $3, adress = $4, adress_number = $5, adress_complement = $6, cep = $7, city = $8
                neighborhood = $9, residense_type = $10, rent_value = $11, building_materials = $12, in_house_bathroom = $13, residents = $14, rooms = $15,
                beds = $16, car = $17, television = $18, refrigerator = $19, microwave = $20, road_type = $21, refering_person = $22, create_log = $23,
                update_log = $24 WHERE id = $1`

	_, err = conn.Exec(sql, &household.ID, &household.Code, &household.Date, &household.Adress, &household.AdressNumber,
		&household.AdressComplement, &household.CEP, &household.City, &household.Neighborhood, &household.ResidenseType,
		&household.RentValue, &household.BuildingMaterials, &household.InHouseBathroom, &household.Residents, &household.Rooms,
		&household.Beds, &household.Car, &household.Television, &household.Refrigerator, &household.Microwave, &household.RoadType,
		&household.ReferingPerson, &household.CreateLog, &household.UpdateLog)
	if err != nil {
		return uuid.Nil, err
	}

	id = household.ID
	return
}
