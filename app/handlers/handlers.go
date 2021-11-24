package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"vietnam-population-server/app/utils"
)

func GetProvinceList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	provinceList := utils.GetProvinceList(db)
	respondJSON(w, http.StatusOK, provinceList)
}

func GetDistrictListByProvinceCode(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["province_code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	districtList := utils.GetDistrictListByProvinceCode(db, key)
	respondJSON(w, http.StatusOK, districtList)
}

func GetWardListByDistrictCode(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["district_code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	wardList := utils.GetWardListByDistrictCode(db, key)
	respondJSON(w, http.StatusOK, wardList)
}
