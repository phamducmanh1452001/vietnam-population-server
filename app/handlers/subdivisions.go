package handlers

import (
	"database/sql"
	"net/http"
	"vietnam-population-server/app/utils"
)

const countryName = "Việt Nam"

func GetProvinceList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	provinceList, population := utils.GetProvinceList(db)
	subDivRes := SubdivisionResponse{
		Area:       countryName,
		Data:       provinceList,
		Population: uint32(population),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}

func GetDistrictListByProvinceCode(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["province_code"]

	if !ok || len(keys[0]) < 1 {
		respondError(w, 400, "URL Param is missing")
		return
	}

	key := keys[0]

	districtList, v1, v2 := utils.GetDistrictListByProvinceCode(db, key)
	if v1 == utils.ErrorFlag {
		respondError(w, 501, v2)
		return
	}
	subDivRes := SubdivisionResponse{
		Area:       v2,
		Data:       districtList,
		Population: uint32(v1),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}

func GetWardListByDistrictCode(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["district_code"]

	if !ok || len(keys[0]) < 1 {
		go respondError(w, 400, "URL Param is missing")
		return
	}

	key := keys[0]

	wardList, v1, v2 := utils.GetWardListByDistrictCode(db, key)
	if v1 == utils.ErrorFlag {
		respondError(w, 501, v2)
		return
	}
	subDivRes := SubdivisionResponse{
		Area:       v2,
		Data:       wardList,
		Population: uint32(v1),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}
