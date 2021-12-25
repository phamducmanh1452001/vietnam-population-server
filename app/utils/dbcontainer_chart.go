package utils

import (
	"database/sql"
	"fmt"
	"log"
)

func GetReligionChart(db *sql.DB, code string) map[string]int {
	data := map[string]int{
		"khong":             getNumberCitizenInReligion(db, 0, code),
		"phat_giao":         getNumberCitizenInReligion(db, 1, code),
		"cong_giao":         getNumberCitizenInReligion(db, 2, code),
		"hoi_giao":          getNumberCitizenInReligion(db, 3, code),
		"tin_lanh":          getNumberCitizenInReligion(db, 4, code),
		"cao_dai":           getNumberCitizenInReligion(db, 5, code),
		"ton_giao_dan_gian": getNumberCitizenInReligion(db, 6, code),
		"hoa_hao":           getNumberCitizenInReligion(db, 7, code),
		"khac":              getNumberCitizenInReligion(db, 8, code),
	}
	return data
}

func GetAgeChart(db *sql.DB, code string) map[string]int {
	data := map[string]int{
		"0-10":  getNumberCitizenInAgeRange(db, 0, 10, code),
		"10-20": getNumberCitizenInAgeRange(db, 10, 20, code),
		"20-30": getNumberCitizenInAgeRange(db, 20, 30, code),
		"30-40": getNumberCitizenInAgeRange(db, 30, 40, code),
		"40-50": getNumberCitizenInAgeRange(db, 40, 50, code),
		"50-60": getNumberCitizenInAgeRange(db, 50, 60, code),
		"60-70": getNumberCitizenInAgeRange(db, 60, 70, code),
		"70-80": getNumberCitizenInAgeRange(db, 70, 80, code),
		"80-90": getNumberCitizenInAgeRange(db, 80, 90, code),
		">=90":  getNumberCitizenInAgeRange(db, 90, 200, code),
	}
	return data
}

func GetGenderChart(db *sql.DB, code string) map[string]int {
	data := map[string]int{
		"nam": getNumberCitzenInGender(db, "M", code),
		"nu":  getNumberCitzenInGender(db, "F", code),
	}
	return data
}

func getNumberCitizenInReligion(db *sql.DB, id int, code string) int {
	number := 0
	codeQuery := ""
	if code != "" {
		codeQuery = fmt.Sprintf("AND %s = '%s'", getCodeName(code), code)
	}
	query := fmt.Sprintf(`SELECT COUNT(religion) FROM citizens 
		WHERE religion_id = %d %s`, id, codeQuery)
	log.Println(query)
	results, _ := db.Query(query)
	if results.Next() {
		results.Scan(&number)
	}
	results.Close()
	return number
}

func getNumberCitizenInAgeRange(db *sql.DB, start int, end int, code string) int {
	number := 0
	codeQuery := ""
	if code != "" {
		codeQuery = fmt.Sprintf("AND %s = '%s'", getCodeName(code), code)
	}
	query := fmt.Sprintf(`SELECT COUNT(age) FROM citizens 
		WHERE age >= %d AND age < %d %s`, start, end, codeQuery)
	log.Println(query)
	results, _ := db.Query(query)
	if results.Next() {
		results.Scan(&number)
	}
	results.Close()
	return number
}

func getNumberCitzenInGender(db *sql.DB, gender string, code string) int {
	number := 0
	codeQuery := ""
	if code != "" {
		codeQuery = fmt.Sprintf("AND %s = '%s'", getCodeName(code), code)
	}
	query := fmt.Sprintf(`SELECT COUNT(gender) FROM citizens 
		WHERE gender = '%s' %s`, gender, codeQuery)
	log.Println(query)
	results, _ := db.Query(query)
	if results.Next() {
		results.Scan(&number)
	}
	results.Close()
	return number
}

func getCodeName(code string) string {
	var codeName string
	switch len(code) {
	case 2:
		codeName = "province_code"
	case 3:
		codeName = "district_code"
	default:
		codeName = "ward_code"
	}
	return codeName
}
