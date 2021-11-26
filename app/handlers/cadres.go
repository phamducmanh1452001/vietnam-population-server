package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"vietnam-population-server/app/utils"
)

func GetLowerCadreListByCode(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)

	if err != nil {
		respondError(w, unauthorizedStatus.number, err.Error())
	}

	var cadreListResponse CadreListResponse
	var cadreResponseArray []CadreResponse

	provinceCodeLen := 2
	districtCodeLen := 3

	var subdivision interface{}
	var population uint32 = 0
	var area string
	code := fmt.Sprintf("%v", claims["code"])

	switch len(code) {
	case provinceCodeLen:
		districtList, v1, v2 := utils.GetDistrictListByProvinceCode(db, code)
		if v1 == utils.ErrorFlag {
			respondError(w, internalErrorStatus.number, v2)
		}
		population = uint32(v1)
		area = v2
		len := len(districtList)

		wg := sync.WaitGroup{}
		for i := 0; i < len; i++ {
			wg.Add(1)
			go func(i int) {
				subdivision = districtList[i]
				cadre, err := utils.GetCadreByCode(db, districtList[i].Code)
				if err != nil {
					respondError(w, notImplementedStatus.number, err.Error())
					wg.Done()
				} else {
					cadreResponse := CadreResponse{
						Name:        cadre.Name.String,
						Code:        cadre.Code,
						Age:         uint8(cadre.Age.Int16),
						Phone:       cadre.Phone.String,
						Email:       cadre.Email.String,
						Permission:  uint8(cadre.Permission),
						Subdivision: subdivision,
					}
					cadreResponseArray = append(cadreResponseArray, cadreResponse)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	case districtCodeLen:
		wardList, v1, v2 := utils.GetWardListByDistrictCode(db, code)
		if v1 == utils.ErrorFlag {
			respondError(w, internalErrorStatus.number, v2)
		}
		population = uint32(v1)
		len := len(wardList)

		wg := sync.WaitGroup{}
		for i := 0; i < len; i++ {
			wg.Add(1)
			go func(i int) {
				subdivision = wardList[i]
				cadre, err := utils.GetCadreByCode(db, wardList[i].Code)
				if err != nil {
					respondError(w, notImplementedStatus.number, err.Error())
				} else {
					cadreResponse := CadreResponse{
						Name:        cadre.Name.String,
						Code:        cadre.Code,
						Age:         uint8(cadre.Age.Int16),
						Phone:       cadre.Phone.String,
						Email:       cadre.Email.String,
						Permission:  uint8(cadre.Permission),
						Subdivision: subdivision,
					}
					cadreResponseArray = append(cadreResponseArray, cadreResponse)
				}
				wg.Done()
			}(i)
		}
	}

	cadreListResponse = CadreListResponse{
		Area:       area,
		Population: population,
		Data:       cadreResponseArray,
	}
	respondJSON(w, http.StatusOK, cadreListResponse)
}
