package handlers

import (
	"database/sql"
	"net/http"
	"vietnam-population-server/app/utils"
)

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		respondError(w, badRequestStatus.number, badRequestStatus.description)
		return
	}

	code := r.PostForm.Get("code")
	password := r.PostForm.Get("password")
	cadre, err := utils.GetCadreByCodeAndPassword(db, code, password)
	if err != nil {
		respondError(w, notFoundStatus.number, err.Error())
		return
	}

	tokenString, err := generateJWT(cadre)
	if err != nil {
		respondError(w, notImplementedStatus.number, notImplementedStatus.description)
		return
	}

	jwtResponse := JwtResponse{Token: tokenString}
	respondJSON(w, http.StatusOK, jwtResponse)
}
