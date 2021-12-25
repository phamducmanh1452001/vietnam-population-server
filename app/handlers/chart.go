package handlers

import (
	"database/sql"
	"net/http"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

func GetReligionChart(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	code, err := getParam(r, "code")
	var json map[string]int
	if err != nil {
		json = utils.GetReligionChart(db, "")
	} else {
		json = utils.GetReligionChart(db, code)
	}
	respondJSON(w, 200, json)
}

func GetAgeChart(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	code, err := getParam(r, "code")
	var json map[string]int
	if err != nil {
		json = utils.GetAgeChart(db, "")
	} else {
		json = utils.GetAgeChart(db, code)
	}
	respondJSON(w, 200, json)
}

func GetGenderChart(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	code, err := getParam(r, "code")
	var json map[string]int
	if err != nil {
		json = utils.GetGenderChart(db, "")
	} else {
		json = utils.GetGenderChart(db, code)
	}
	respondJSON(w, 200, json)
}
