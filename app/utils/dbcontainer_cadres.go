package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
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

func GetCadreListBySuperCode(db *sql.DB, superCode string, page int, limit int, key string) ([]cadre.Cadre, error) {
	cadreList := []cadre.Cadre{}

	table := "cadres"
	fields := "code, name, password, age, phone, email"
	offset := (page - 1) * limit
	codeSearch := " true "
	nameSearch := " true "
	words := getSearchKeyArray(key)
	for _, word := range words {
		if word == "" {
			continue
		}
		lowerWord := strings.ToLower(word)
		codeSearch += " AND code LIKE '%" + lowerWord + "%' "
		nameSearch += " AND LOWER(name) LIKE BINARY '%" + lowerWord + "%' "
	}
	condition := fmt.Sprintf("WHERE super_code = '%s' and ((%s) or (%s)) LIMIT %d OFFSET %d",
		superCode, nameSearch, codeSearch, limit, offset)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		log.Println("Error: ", err.Error())
		return cadreList, err
	}

	var cadre cadre.Cadre

	for results.Next() {
		err := results.Scan(&cadre.Code, &cadre.Name, &cadre.Password, &cadre.Age, &cadre.Phone, &cadre.Email)
		if err != nil {
			return cadreList, errors.New("cannot scan result from database")
		}
		cadre.Permission, err = GetCadrePermissionByCode(db, cadre.Code)
		if err != nil {
			return cadreList, errors.New("cannot get permission result from database")
		}
		cadreList = append(cadreList, cadre)
	}
	results.Close()

	return cadreList, nil
}

func GetCadreByCode(db *sql.DB, code string) (cadre.Cadre, error) {
	cadre := cadre.Cadre{}

	table := "cadres"
	fields := "code, name, password, age, phone, email, permission, super_code"
	condition := fmt.Sprintf("WHERE code = '%s'", code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return cadre, err
	}

	if results.Next() {
		err := results.Scan(&cadre.Code, &cadre.Name, &cadre.Password, &cadre.Age,
			&cadre.Phone, &cadre.Email, &cadre.Permission, &cadre.SuperCode)
		if err != nil {
			return cadre, errors.New("cannot scan result from database")
		}
	}
	results.Close()

	return cadre, nil
}

func GetCadrePermissionByCode(db *sql.DB, code string) (int, error) {
	cadre, err := GetCadreByCode(db, code)
	if err != nil {
		return 0, err
	}
	if cadre.Permission == 0 {
		return 0, nil
	}
	if cadre.SuperCode == "" {
		return 1, nil
	}
	return GetCadrePermissionByCode(db, cadre.SuperCode)
}

func ChangeCadrePermisson(db *sql.DB, code string, permission int) error {
	table := "cadres"
	fields := "permission"
	condition := fmt.Sprintf("WHERE code = '%s'", code)
	query := fmt.Sprintf("UPDATE %s SET %s = %d %s", table, fields, permission, condition)

	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	return nil
}
