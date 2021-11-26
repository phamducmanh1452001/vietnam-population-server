package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	cadre "vietnam-population-server/app/models/cadre"

	_ "github.com/go-sql-driver/mysql"
)

func GetCadreByCodeAndPassword(db *sql.DB, code string, password string) (cadre.Cadre, error) {
	cadre := cadre.Cadre{}

	table := "cadres"
	fields := "code, name, password, age, phone, email"
	condition := fmt.Sprintf("WHERE code = '%s' AND password = '%s'", code, password)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return cadre, err
	}

	if results.Next() {
		err := results.Scan(&cadre.Code, &cadre.Name, &cadre.Password, &cadre.Age, &cadre.Phone, &cadre.Email)
		if err != nil {
			return cadre, errors.New("cannot scan result from database")
		}
		return cadre, nil
	}

	return cadre, errors.New("code or password was wrong")
}

func GetCadreByCode(db *sql.DB, code string) (cadre.Cadre, error) {
	cadre := cadre.Cadre{}

	table := "cadres"
	fields := "code, name, password, age, phone, email, permission"
	condition := fmt.Sprintf("WHERE code = '%s'", code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return cadre, err
	}

	if results.Next() {
		err := results.Scan(&cadre.Code, &cadre.Name, &cadre.Password, &cadre.Age, &cadre.Phone, &cadre.Email, &cadre.Permission)
		if err != nil {
			return cadre, errors.New("cannot scan result from database")
		}
	}

	return cadre, nil
}