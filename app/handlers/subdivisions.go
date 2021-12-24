package handlers

import (
	"database/sql"
	"net/http"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

const countryName = "Việt Nam"

func GetProvinceList(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	page, limit := getPageAndLimit(r)
	searchKey, _ := getParam(r, "key")

	provinceList, population, amount := utils.GetProvinceList(db, page, limit, searchKey)
	var areaSize float64 = 0
	for _, v := range provinceList {
		areaSize = areaSize + v.Area
	}
	subDivRes := SubdivisionResponse{
		AreaSize:   areaSize,
		Area:       countryName,
		Amount:     amount,
		Data:       provinceList,
		Population: uint32(population),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}

func GetDistrictListByProvinceCode(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	provinceCode, err := getParam(r, "province_code")
	searchKey, _ := getParam(r, "key")

	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	page, limit := getPageAndLimit(r)

	districtList, v1, v2, amount := utils.GetDistrictListByProvinceCode(db, provinceCode, page, limit, searchKey)
	if v1 == utils.ErrorFlag {
		respondError(w, 501, v2)
		return
	}
	var areaSize float64 = 0
	for _, v := range districtList {
		areaSize = areaSize + v.Area
	}
	subDivRes := SubdivisionResponse{
		AreaSize:   areaSize,
		Area:       v2,
		Amount:     amount,
		Data:       districtList,
		Population: uint32(v1),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}

func GetWardListByDistrictCode(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	districtCode, err := getParam(r, "district_code")
	searchKey, _ := getParam(r, "key")

	if err != nil {
		respondError(w, 400, err.Error())
		return
	}

	page, limit := getPageAndLimit(r)
	wardList, v1, v2, amount := utils.GetWardListByDistrictCode(db, districtCode, page, limit, searchKey)
	if v1 == utils.ErrorFlag {
		respondError(w, 501, v2)
		return
	}
	var areaSize float64 = 0
	for _, v := range wardList {
		areaSize = areaSize + v.Area
	}
	subDivRes := SubdivisionResponse{
		AreaSize:   areaSize,
		Area:       v2,
		Amount:     amount,
		Data:       wardList,
		Population: uint32(v1),
	}
	respondJSON(w, http.StatusOK, subDivRes)
}
