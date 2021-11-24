package utils

import (
	"database/sql"
	"fmt"
	"log"
	subDivs "vietnam-population-server/app/models/subdivisions"

	_ "github.com/go-sql-driver/mysql"
)

func GetProvinceList(db *sql.DB) []subDivs.Province {
	provinceList := []subDivs.Province{}

	table := "provinces"
	fields := "code, name"
	query := fmt.Sprintf("SELECT %s FROM %s", fields, table)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return provinceList
	}

	var province subDivs.Province
	for results.Next() {
		err = results.Scan(&province.Code, &province.Name)
		if err != nil {
			log.Println(err.Error())
			return provinceList
		}

		provinceList = append(provinceList, province)
	}

	return provinceList
}

func GetDistrictListByProvinceCode(db *sql.DB, province_code string) []subDivs.District {
	districtList := []subDivs.District{}

	table := "districts"
	fields := "code, name, province_code"
	condition := fmt.Sprintf("WHERE province_code = '%s'", province_code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return districtList
	}

	var district subDivs.District
	for results.Next() {
		err = results.Scan(&district.Code, &district.Name, &district.ProvinceCode)
		if err != nil {
			log.Println(err.Error())
			return districtList
		}

		districtList = append(districtList, district)
	}

	return districtList
}

func GetWardListByDistrictCode(db *sql.DB, district_code string) []subDivs.Ward {
	wardList := []subDivs.Ward{}

	table := "wards"
	fields := "code, name, district_code"
	condition := fmt.Sprintf("WHERE district_code = '%s'", district_code)
	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, table, condition)

	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return wardList
	}

	var ward subDivs.Ward
	for results.Next() {
		err = results.Scan(&ward.Code, &ward.Name, &ward.DistrictCode)
		if err != nil {
			log.Println(err.Error())
			return wardList
		}

		wardList = append(wardList, ward)
	}

	return wardList
}
