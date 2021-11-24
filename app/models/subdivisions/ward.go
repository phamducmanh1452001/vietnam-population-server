package models

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Ward struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	DistrictCode string `json:"district_code"`
}

func GetWardListByDistrictCode(code string) []Ward {
	wardList := []Ward{}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := dir + "/app/models/subdivisions/datasets/ward.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newWard Ward
	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), ",")

		districtCode := strings[2]
		if districtCode == code {
			newWard = Ward{Name: strings[0], Code: strings[1], DistrictCode: districtCode}
			wardList = append(wardList, newWard)
		}
	}

	return wardList
}
