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

	var codeName string
	switch len(cadreCode) {
	case 2:
		codeName = "province_code"
	case 3:
		codeName = "district_code"
	default:
		codeName = "ward_code"
	}
	condition := fmt.Sprintf("WHERE %s = '%s'", codeName, cadreCode)
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
	results.Close()

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

func AddCitizen(db *sql.DB, citizen citizen.Citizen, wardCode string) error {
	ward, err := GetWardByCode(db, wardCode)
	if err != nil {
		return err
	}
	district, err := GetDistrictByCode(db, ward.SuperCode)
	if err != nil {
		return err
	}

	table := "citizens"
	fields := `code, first_name, middle_name, last_name, gender, date_of_birth, age,
		weight, date_of_joining, religion, avatar, collaborator_name, collaborator_phone, 
		ward_code, district_code, province_code`
	values := fmt.Sprintf(`'%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', '%s', 
		'%s', '%s', '%s', '%s', '%s', '%s'`,
		citizen.Code, citizen.FirstName, citizen.MiddleName, citizen.LastName,
		citizen.Gender, citizen.DateOfBirth, citizen.Age, citizen.Weight, citizen.DateOfJoining,
		citizen.Religion, citizen.Avatar, citizen.CollaboratorName, citizen.CollaboratorPhone,
		wardCode, district.Code, district.SuperCode)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fields, values)
	log.Println(query)
	_, err = db.Query(query)
	return err
}
