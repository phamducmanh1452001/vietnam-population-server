package models

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Province struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetProvinceList() []Province {
	provinceList := []Province{}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := dir + "/app/models/subdivisions/datasets/province.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newProvince Province
	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), ",")
		newProvince = Province{Name: strings[0], Code: strings[1]}
		provinceList = append(provinceList, newProvince)
	}

	return provinceList
}
