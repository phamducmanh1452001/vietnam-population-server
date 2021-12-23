package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"vietnam-population-server/app/router"
	"vietnam-population-server/app/utils"
)

func GetCitizenList(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		respondError(w, unauthorizedStatus.number, err.Error())
		return
	}

	page, limit := getPageAndLimit(r)
	searchKey, _ := getParam(r, "key")

	cadreCode := fmt.Sprintf("%v", claims["code"])

	citizenList, amount, err := utils.GetCitizenListByCadreCode(db, cadreCode, page, limit, searchKey)
	if err != nil {
		respondError(w, internalErrorStatus.number, err.Error())
		return
	}

	citizenListResponse := CitizenListResponse{
		Amount: amount,
		Data:   citizenList,
	}

	respondJSON(w, http.StatusOK, citizenListResponse)
}
