package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	citizen "vietnam-population-server/app/models/citizen"
)

func GetCitizenListByCadreCode(db *sql.DB, cadreCode string, page int, limit int, key string) ([]citizen.Citizen, int, error) {
	citizenList := []citizen.Citizen{}
	amount := 0

	table := "citizens"
	fields := `code, first_name, middle_name, last_name, gender, date_of_birth, age,
		date_of_joining, religion_id, avatar, collaborator_name, collaborator_phone, temporary_address, major`
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
		log.Println("Error1: ", err.Error())
		return citizenList, amount, err
	}

	citizen := citizen.Citizen{}

	for results.Next() {
		err = results.Scan(&citizen.Code, &citizen.FirstName, &citizen.MiddleName, &citizen.LastName,
			&citizen.Gender, &citizen.DateOfBirth, &citizen.Age,
			&citizen.DateOfJoining, &citizen.ReligionId, &citizen.Avatar, &citizen.CollaboratorName,
			&citizen.CollaboratorPhone, &citizen.Major, &citizen.Major)
		if err != nil {
			return citizenList, 0, errors.New("cannot scan result from database")
		}
		citizenList = append(citizenList, citizen)
	}
	results.Close()

	query = fmt.Sprintf("SELECT COUNT(code) FROM %s %s", table, condition)
	log.Println(query)
	results, err = db.Query(query)
	if err != nil {
		log.Println("Error2: ", err.Error())
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

	citizen.Age = getAgeFromBirthDay(citizen.DateOfBirth)

	table := "citizens"
	fields := `code, first_name, middle_name, last_name, gender, date_of_birth, age,
		date_of_joining, religion_id, avatar, collaborator_name, collaborator_phone, 
		ward_code, district_code, province_code, temporary_address, major`
	values := fmt.Sprintf(`'%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', %d, 
		'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'`,
		citizen.Code, citizen.FirstName, citizen.MiddleName, citizen.LastName,
		citizen.Gender, citizen.DateOfBirth, citizen.Age, citizen.DateOfJoining,
		citizen.ReligionId, citizen.Avatar, citizen.CollaboratorName, citizen.CollaboratorPhone,
		wardCode, district.Code, district.SuperCode, citizen.TemporaryAddress, citizen.Major)
	query := ""
	if !isCitizenExisted(db, citizen.Code) {
		query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fields, values)
	} else {
		updateStr := fmt.Sprintf(`code = '%s', first_name = '%s', middle_name = '%s', last_name = '%s', 
			gender = '%s', date_of_birth = '%s', age = %d, date_of_joining = '%s', religion_id = %d, 
			avatar = '%s', collaborator_name = '%s', collaborator_phone = '%s', ward_code = '%s', 
			district_code = '%s', province_code = '%s', temporary_address = '%s', major = '%s'`,
			citizen.Code, citizen.FirstName, citizen.MiddleName, citizen.LastName,
			citizen.Gender, citizen.DateOfBirth, citizen.Age, citizen.DateOfJoining,
			citizen.ReligionId, citizen.Avatar, citizen.CollaboratorName, citizen.CollaboratorPhone,
			wardCode, district.Code, district.SuperCode, citizen.TemporaryAddress, citizen.Major)
		query = fmt.Sprintf("UPDATE citizens SET %s WHERE code = %s", updateStr, citizen.Code)
	}

	log.Println(query)
	_, err = db.Query(query)
	return err
}

func getAgeFromBirthDay(date string) int {
	btime, _ := time.Parse("2006-01-02", date)
	ctime := time.Now()
	years := ctime.Sub(btime).Hours() / 24 / 365
	return int(years)
}

func isCitizenExisted(db *sql.DB, code string) bool {
	query := fmt.Sprintf("SELECT COUNT(code) FROM citizens WHERE code = '%s'", code)
	log.Println(query)
	results, _ := db.Query(query)

	count := 0
	if results.Next() {
		results.Scan(&count)
	}

	results.Close()
	return count > 0
}
