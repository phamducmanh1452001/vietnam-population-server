package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	citizen "vietnam-population-server/app/models/citizen"
)

func GetCitizenListByCadreCode(db *sql.DB, cadreCode string, page int, limit int, key string) ([]citizen.Citizen, int, error) {
	citizenList := []citizen.Citizen{}
	amount := 0

	table := "citizens"
	fields := `code, first_name, middle_name, last_name, gender,
		date_of_birth, age, weight, date_of_joining, religion, avatar`
	offset := (page - 1) * limit

	var code string
	switch len(cadreCode) {
	case 2:
		code = "province_code"
	case 3:
		code = "district_code"
	default:
		code = "ward_code"
	}
	condition := fmt.Sprintf("WHERE %s = '%s'", code, cadreCode)
	query := fmt.Sprintf("SELECT %s FROM %s %s LIMIT %d OFFSET %d", fields, table, condition, limit, offset)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		log.Println("Error: ", err.Error())
		return citizenList, amount, err
	}

	citizen := citizen.Citizen{}
	var religion sql.NullString
	var middleName sql.NullString

	for results.Next() {
		err = results.Scan(&citizen.Code, &citizen.FirstName, &middleName, &citizen.LastName,
			&citizen.Gender, &citizen.DateOfBirth, &citizen.Age, &citizen.Weight,
			&citizen.DateOfJoining, &religion, &citizen.Avatar)
		if err != nil {
			return citizenList, 0, errors.New("cannot scan result from database")
		}
		citizen.Religion = religion.String
		citizen.MiddleName = middleName.String

		citizenList = append(citizenList, citizen)
	}

	query = fmt.Sprintf("SELECT COUNT(code) FROM %s %s", table, condition)
	log.Println(query)
	results, err = db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return citizenList, amount, err
	}

	if results.Next() {
		results.Scan(&amount)
	}

	results.Close()

	return citizenList, amount, nil
}
