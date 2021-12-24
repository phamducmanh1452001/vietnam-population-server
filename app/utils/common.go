package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	subDivs "vietnam-population-server/app/models/subdivisions"
)

func getSearchKeyArray(key string) []string {
	words := strings.Split(key, " ")

	// length := len(key)
	// for i := 0; i < length-3; i++ {
	// 	words = append(words, strings.Trim(key[i:i+3], " "))
	// }
	return words
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

func getPopulationBySubdivisionCode(db *sql.DB, code string) (uint32, error) {
	var population uint32 = 0

	table := "citizens"
	var codeName string
	switch len(code) {
	case 2:
		codeName = "province_code"
	case 3:
		codeName = "district_code"
	default:
		codeName = "ward_code"
	}
	condition := fmt.Sprintf("WHERE %s = %s", codeName, code)
	query := fmt.Sprintf("SELECT COUNT(code) FROM %s %s", table, condition)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return population, err
	}
	if results.Next() {
		results.Scan(&population)
	}
	results.Close()

	return population, nil
}
