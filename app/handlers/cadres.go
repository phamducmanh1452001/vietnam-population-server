package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

func GetLowerCadreListByCode(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	page, limit := getPageAndLimit(r)
	searchKey, _ := getParam(r, "key")

	if err != nil {
		respondError(w, unauthorizedStatus.number, err.Error())
		return
	}

	var cadreListResponse CadreListResponse
	var cadreResponseArray []CadreResponse

	provinceCodeLen := 2
	districtCodeLen := 3
	adminCodeLen := 5

	var subdivision interface{}
	var population uint32 = 0
	var area string
	var amount int = 0
	code := fmt.Sprintf("%v", claims["code"])

	cadreList, err := utils.GetCadreListBySuperCode(db, code, page, limit, searchKey)

	if err != nil {
		respondError(w, internalErrorStatus.number, err.Error())
		return
	}

	switch len(code) {
	case provinceCodeLen:
		districtList, v1, v2, cnt := utils.GetDistrictListByProvinceCode(db, code, page, limit, searchKey)
		amount = cnt
		if v1 == utils.ErrorFlag {
			respondError(w, internalErrorStatus.number, "Database error1")
			return
		}
		population = uint32(v1)
		area = v2

		if len(cadreList) != len(districtList) {
			fmt.Println(len(cadreList), " ", len(districtList))
			respondError(w, internalErrorStatus.number, "Database error2")
			return
		}

		for i, cadre := range cadreList {
			subdivision = districtList[i]
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

	case districtCodeLen:
		wardList, v1, v2, cnt := utils.GetWardListByDistrictCode(db, code, page, limit, searchKey)
		amount = cnt
		if v1 == utils.ErrorFlag {
			respondError(w, internalErrorStatus.number, v2)
			return
		}
		population = uint32(v1)
		area = v2

		if len(cadreList) != len(wardList) {
			respondError(w, internalErrorStatus.number, v2)
			return
		}

		for i, cadre := range cadreList {
			subdivision = wardList[i]
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
	case adminCodeLen:
		provinceList, v1, cnt := utils.GetProvinceList(db, page, limit, searchKey)
		amount = cnt
		if v1 == utils.ErrorFlag {
			respondError(w, internalErrorStatus.number, "Database error1")
			return
		}
		population = uint32(v1)
		area = "Việt Nam"

		if len(cadreList) != len(provinceList) {
			fmt.Println(len(cadreList), " ", len(provinceList))
			respondError(w, internalErrorStatus.number, "Database error2")
			return
		}

		for i, cadre := range cadreList {
			subdivision = provinceList[i]
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
	}

	permission, err := utils.GetCadrePermissionByCode(db, code)
	if err != nil {
		respondError(w, internalErrorStatus.number, err.Error()+"flagqweu")
	}
	cadreListResponse = CadreListResponse{
		Permission: permission,
		Area:       area,
		Amount:     amount,
		Population: population,
		Data:       cadreResponseArray,
	}
	respondJSON(w, http.StatusOK, cadreListResponse)
}
