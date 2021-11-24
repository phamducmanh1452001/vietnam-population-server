package models

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type District struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	ProvinceCode string `json:"province_code"`
}

func GetDistrictListByProvinceCode(code string) []District {
	districtList := []District{}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := dir + "/app/models/subdivisions/datasets/district.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newDistrict District
	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), ",")

		provinceCode := strings[2]
		if provinceCode == code {
			newDistrict = District{Name: strings[0], Code: strings[1], ProvinceCode: provinceCode}
			districtList = append(districtList, newDistrict)
		}
	}

	return districtList
}
