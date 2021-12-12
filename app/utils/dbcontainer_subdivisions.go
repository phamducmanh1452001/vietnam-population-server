package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	subDivs "vietnam-population-server/app/models/subdivisions"

	_ "github.com/go-sql-driver/mysql"
)

const ErrorFlag uint32 = 12344321
const queryError = "Query SQL Error"

func GetProvinceList(db *sql.DB, page int, limit int, key string) ([]subDivs.Province, uint32, int) {
	provinceList := []subDivs.Province{}
	var population uint32 = 0
	var err error
	var results *sql.Rows
	var amount int = 0

	table := "provinces"
	fields := "code, name, population"
	offset := (page - 1) * limit
	words := getSearchKeyArray(key)
	codeSearch := " true "
	nameSearch := " true "
	for _, word := range words {
		if word == "" {
			continue
		}

		lowerWord := strings.ToLower(word)
		codeSearch += " AND code LIKE '%" + lowerWord + "%' "
		nameSearch += " AND LOWER(name) LIKE BINARY '%" + lowerWord + "%' "
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE (%s) or (%s) LIMIT %d OFFSET %d",
		fields, table, codeSearch, nameSearch, limit, offset)

	log.Println(query)
	results, err = db.Query(query)

	if err != nil {
		log.Println("Error: ", err.Error())
		return provinceList, ErrorFlag, amount
	}

	var province subDivs.Province
	for results.Next() {
		err = results.Scan(&province.Code, &province.Name, &province.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
		} else {
			population += province.Population
			province.Level = getLevelFromSubdivisionName(province.Name, province.Code)
			provinceList = append(provinceList, province)
		}
	}

	amountQuery := fmt.Sprintf("SELECT COUNT(code) FROM %s WHERE (%s) or (%s)",
		table, codeSearch, nameSearch)
	log.Println(amountQuery)
	results, err = db.Query(amountQuery)

	if err != nil {
		log.Println("Error: ", err.Error())
		return provinceList, ErrorFlag, amount
	}

	if results.Next() {
		err = results.Scan(&amount)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
	}

	return provinceList, population, amount
}

func GetDistrictListByProvinceCode(db *sql.DB, provinceCode string, page int, limit int, key string) ([]subDivs.District, uint32, string, int) {
	districtList := []subDivs.District{}
	var population uint32 = 0
	var amount = 0

	table := "districts"
	fields := "code, name, super_code, population"
	offset := (page - 1) * limit
	words := getSearchKeyArray(key)
	codeSearch := " true "
	nameSearch := " true "
	for _, word := range words {
		if word == "" {
			continue
		}
		lowerWord := strings.ToLower(word)
		codeSearch += " AND code LIKE '%" + lowerWord + "%' "
		nameSearch += " AND LOWER(name) LIKE BINARY '%" + lowerWord + "%' "
	}
	condition := fmt.Sprintf("WHERE super_code = '%s' and ((%s) or (%s)) LIMIT %d OFFSET %d",
		provinceCode, codeSearch, nameSearch, limit, offset)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		log.Println("Error: ", err.Error())
		return districtList, ErrorFlag, queryError, amount
	}
	var district subDivs.District

	for results.Next() {
		err = results.Scan(&district.Code, &district.Name, &district.SuperCode, &district.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
		} else {
			population += district.Population
			district.Level = getLevelFromSubdivisionName(district.Name, district.Code)
			districtList = append(districtList, district)
		}
	}

	amountQuery := fmt.Sprintf("SELECT COUNT(code) FROM %s %s",
		table, condition)
	log.Println(amountQuery)
	results, err = db.Query(amountQuery)
	if err != nil {
		log.Println("Error: ", err.Error())
		return districtList, ErrorFlag, queryError, amount
	}

	if results.Next() {
		err = results.Scan(&amount)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
	}

	province, err := GetProvinceByCode(db, provinceCode)
	if err != nil {
		log.Println("Error: ", err.Error())
		return districtList, ErrorFlag, queryError, amount
	}

	return districtList, population, province.Name, amount
}

func GetWardListByDistrictCode(db *sql.DB, districtCode string, page int, limit int, key string) ([]subDivs.Ward, uint32, string, int) {
	wardList := []subDivs.Ward{}
	var population uint32 = 0
	var amount int = 0

	table := "wards"
	fields := "code, name, super_code, population"
	offset := (page - 1) * limit
	words := getSearchKeyArray(key)
	codeSearch := " true "
	nameSearch := " true "
	for _, word := range words {
		if word == "" {
			continue
		}
		lowerWord := strings.ToLower(word)
		codeSearch += " AND code LIKE '%" + lowerWord + "%' "
		nameSearch += " AND LOWER(name) LIKE BINARY '%" + lowerWord + "%' "
	}
	condition := fmt.Sprintf("WHERE super_code = '%s' and ((%s) or (%s)) LIMIT %d OFFSET %d", districtCode, codeSearch, nameSearch, limit, offset)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		log.Println("Error: ", err.Error())
		return wardList, ErrorFlag, queryError, amount
	}

	var ward subDivs.Ward
	for results.Next() {
		err = results.Scan(&ward.Code, &ward.Name, &ward.SuperCode, &ward.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
		} else {
			population += ward.Population
			ward.Level = getLevelFromSubdivisionName(ward.Name, ward.Code)
			wardList = append(wardList, ward)
		}
	}

	amountQuery := fmt.Sprintf("SELECT COUNT(code) FROM %s %s",
		table, condition)
	log.Println(amountQuery)
	results, err = db.Query(amountQuery)
	if err != nil {
		log.Println("Error: ", err.Error())
		return wardList, ErrorFlag, queryError, amount
	}

	if results.Next() {
		err = results.Scan(&amount)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
	}

	district, err := GetDistrictByCode(db, districtCode)
	if err != nil {
		log.Println("Error: ", err.Error())
		return wardList, ErrorFlag, queryError, amount
	}

	return wardList, population, district.Name, amount
}

func GetProvinceByCode(db *sql.DB, code string) (subDivs.Province, error) {
	province := subDivs.Province{}

	table := "provinces"
	fields := "code, name, population"
	condition := fmt.Sprintf("WHERE	code = '%s'", code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		return province, err
	}

	if results.Next() {
		err = results.Scan(&province.Code, &province.Name, &province.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
			return province, err
		}

		province.Level = getLevelFromSubdivisionName(province.Name, province.Code)
	}

	return province, nil
}

func GetDistrictByCode(db *sql.DB, code string) (subDivs.District, error) {
	district := subDivs.District{}

	table := "districts"
	fields := "code, name, population, super_code"
	condition := fmt.Sprintf("WHERE	code = '%s'", code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		return district, err
	}

	if results.Next() {
		err = results.Scan(&district.Code, &district.Name, &district.Population, &district.SuperCode)
		if err != nil {
			log.Println("Error: ", err.Error())
			return district, err
		}

		district.Level = getLevelFromSubdivisionName(district.Name, district.Code)
	}

	return district, nil
}

func GetWardByCode(db *sql.DB, code string) (subDivs.Ward, error) {
	ward := subDivs.Ward{}

	table := "wards"
	fields := "code, name, population, super_code"
	condition := fmt.Sprintf("WHERE	code = '%s'", code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)

	if err != nil {
		return ward, err
	}

	if results.Next() {
		err = results.Scan(&ward.Code, &ward.Name, &ward.Population, &ward.SuperCode)
		if err != nil {
			log.Println("Error: ", err.Error())
			return ward, err
		}

		ward.Level = getLevelFromSubdivisionName(ward.Name, ward.Code)
	}

	return ward, nil
}

func getLevelFromSubdivisionName(name string, code string) subDivs.SubdivisionLevel {
	levelNames := []string{"Tỉnh", "Thành phố", "Quận", "Huyện", "Thị xã", "Thị trấn", "Xã", "Phường"}
	levelNumber := map[int]uint8{5: 1, 3: 2, 2: 1}

	level := subDivs.SubdivisionLevel{Name: "", Number: 0}
	for _, v := range levelNames {
		if strings.HasPrefix(name, v) {
			level.Name = v
			level.Number = levelNumber[len(code)]
			return level
		}
	}
	return level
}
