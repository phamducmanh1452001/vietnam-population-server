package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

func ChangeCadrePermisson(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		respondError(w, unauthorizedStatus.number, err.Error())
		return
	}

	if err := r.ParseForm(); err != nil {
		respondError(w, badRequestStatus.number, badRequestStatus.description)
		return
	}

	permissionParam, err := strconv.Atoi(r.PostForm.Get("permission"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid type")
		return
	}
	lowerCode := r.PostForm.Get("code")

	code := fmt.Sprintf("%v", claims["code"])

	cadre, err := utils.GetCadreByCode(db, lowerCode)
	if err != nil {
		respondError(w, http.StatusNotImplemented, "Get data cadre error")
		return
	}

	permission, err := utils.GetCadrePermissionByCode(db, code)
	if err != nil {
		respondError(w, http.StatusNotImplemented, "Database error while change permission")
		return
	}

	if permission == 0 {
		respondJSON(w, http.StatusOK, map[string]string{
			"message": "You was denied by higher cadre",
		})
		return
	}

	if code == cadre.SuperCode {
		err := utils.ChangeCadrePermisson(db, lowerCode, permissionParam)
		if err != nil {
			respondError(w, http.StatusNotImplemented, "Database error while change permission")
			return
		}
	} else {
		respondError(w, http.StatusNotImplemented, "You cannot change permission of this cadre")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{
		"message": "changed permission successfully",
	})
}
