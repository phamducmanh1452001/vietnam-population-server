package handlers

import (
	"log"
	"net/http"
	subDivs "vietnam-population-server/app/models/subdivisions"
)

func GetProvinceList(w http.ResponseWriter, r *http.Request) {
	provinceList := subDivs.GetProvinceList()
	respondJSON(w, http.StatusOK, provinceList)
}

func GetDistrictListByProvinceCode(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["province_code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	districtList := subDivs.GetDistrictListByProvinceCode(key)
	respondJSON(w, http.StatusOK, districtList)
}

func GetWardListByDistrictCode(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["district_code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	wardList := subDivs.GetWardListByDistrictCode(key)
	respondJSON(w, http.StatusOK, wardList)
}
