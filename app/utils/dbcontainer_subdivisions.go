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

func GetProvinceList(db *sql.DB) ([]subDivs.Province, uint32) {
	provinceList := []subDivs.Province{}
	var population uint32 = 0

	table := "provinces"
	fields := "code, name, population"
	query := fmt.Sprintf("SELECT %s FROM %s", fields, table)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error28: ", err.Error())
		return provinceList, ErrorFlag
	}

	var province subDivs.Province
	for results.Next() {
		err = results.Scan(&province.Code, &province.Name, &province.Population)
		if err != nil {
			log.Println("Error39: ", err.Error())
		} else {
			population += province.Population
			province.Level = getLevelFromSubdivisionName(province.Name)
			provinceList = append(provinceList, province)
		}
	}

	return provinceList, population
}

func GetDistrictListByProvinceCode(db *sql.DB, provinceCode string) ([]subDivs.District, uint32, string) {
	districtList := []subDivs.District{}
	var population uint32 = 0

	table := "districts"
	fields := "code, name, super_code, population"
	condition := fmt.Sprintf("WHERE super_code = '%s'", provinceCode)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return districtList, ErrorFlag, queryError
	}
	var district subDivs.District

	for results.Next() {
		err = results.Scan(&district.Code, &district.Name, &district.SuperCode, &district.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
		} else {
			population += district.Population
			district.Level = getLevelFromSubdivisionName(district.Name)
			districtList = append(districtList, district)
		}
	}

	table = "provinces"
	fields = "name"
	condition = fmt.Sprintf("WHERE code = '%s'", provinceCode)
	query = fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)
	results, err = db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return districtList, ErrorFlag, ""
	}

	var area string
	if results.Next() {
		err = results.Scan(&area)
		if err != nil {
			return districtList, ErrorFlag, ""
		}
	}

	return districtList, population, area
}

func GetWardListByDistrictCode(db *sql.DB, districtCode string) ([]subDivs.Ward, uint32, string) {
	wardList := []subDivs.Ward{}
	var population uint32 = 0

	table := "wards"
	fields := "code, name, super_code, population"
	condition := fmt.Sprintf("WHERE super_code = '%s'", districtCode)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return wardList, ErrorFlag, queryError
	}

	var ward subDivs.Ward
	for results.Next() {
		err = results.Scan(&ward.Code, &ward.Name, &ward.SuperCode, &ward.Population)
		if err != nil {
			log.Println("Error: ", err.Error())
		} else {
			population += ward.Population
			ward.Level = getLevelFromSubdivisionName(ward.Name)
			wardList = append(wardList, ward)
		}
	}

	table = "districts"
	fields = "name"
	condition = fmt.Sprintf("WHERE code = '%s'", districtCode)
	query = fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)
	results, err = db.Query(query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return wardList, ErrorFlag, queryError
	}
	var area string
	if results.Next() {
		err = results.Scan(&area)
		if err != nil {
			return wardList, ErrorFlag, queryError
		}

	}
	return wardList, population, area
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

		province.Level = getLevelFromSubdivisionName(province.Name)
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

		district.Level = getLevelFromSubdivisionName(district.Name)
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

		ward.Level = getLevelFromSubdivisionName(ward.Name)
	}

	return ward, nil
}

func getLevelFromSubdivisionName(name string) subDivs.SubdivisionLevel {
	levelMap := make(map[string]uint8)
	levelMap["Tỉnh"] = 3
	levelMap["Thành phố"] = 3
	levelMap["Quận"] = 2
	levelMap["Huyện"] = 2
	levelMap["Thị Xã"] = 2
	levelMap["Xã"] = 1
	levelMap["Thị Trấn"] = 1

	level := subDivs.SubdivisionLevel{Name: "", Number: 0}
	for key, value := range levelMap {
		if strings.HasPrefix(name, key) {
			level.Name = key
			level.Number = value
			return level
		}
	}
	return level
}
