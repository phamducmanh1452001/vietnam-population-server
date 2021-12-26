package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

func Login(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		respondError(w, badRequestStatus.number, badRequestStatus.description)
		return
	}

	code := r.PostForm.Get("code")
	password := r.PostForm.Get("password")
	if code == "" || password == "" {
		respondError(w, badRequestStatus.number, badRequestStatus.description)
		return
	}

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
	permission, err := utils.GetCadrePermissionByCode(db, code)
	if err != nil {
		respondError(w, notImplementedStatus.number, "Error when get permission")
		return
	}
	jwtResponse := JwtResponse{Token: tokenString, Permission: permission}
	respondJSON(w, http.StatusOK, jwtResponse)
}

func Logout(sb *sql.DB, w *router.ResponseWriter, r *http.Request) {
	headerParam := "Authorization"
	if r.Header[headerParam] != nil {
		authString := r.Header[headerParam][0]
		tokenStrings := strings.Split(authString, " ")
		if len(tokenStrings) <= 1 {
			respondError(w, badRequestStatus.number, "Missing Brearer"+headerParam)
			return
		}
		tokenString := tokenStrings[1]
		blackTokenList[tokenString] = 1
		respondJSON(w, 200, map[string]string{"message": "logout success"})
		go func(tokenString string) {
			time.Sleep(expiredTime)
			blackTokenList[tokenString] = 0
		}(tokenString)
	} else {
		respondError(w, badRequestStatus.number, "Missing "+headerParam)
	}
}
