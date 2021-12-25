package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	models "vietnam-population-server/app/models/citizen"
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

func AddCitizen(db *sql.DB, w *router.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		respondError(w, unauthorizedStatus.number, err.Error())
		return
	}

	code := fmt.Sprintf("%v", claims["code"])
	isWardCadre := len(code) == 5

	if !isWardCadre {
		errStr := "Only ward cadre can add new citizen"
		respondError(w, internalErrorStatus.number, errStr)
		return
	}

	citizen, err := citizenFromPostForm(r)
	fmt.Println("Citizen: ", citizen)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.AddCitizen(db, citizen, code)
	if err != nil {
		errString := err.Error()
		respondError(w, internalErrorStatus.number, errString)
		return
	}

	respondJSON(w, 200, citizen)
}

func citizenFromPostForm(r *http.Request) (models.Citizen, error) {
	citizen := models.Citizen{}
	var err error
	var errString string = ""

	if err := r.ParseForm(); err != nil {
		return citizen, err
	}

	form := r.PostForm
	citizen.Code = form.Get("code")
	citizen.FirstName = form.Get("first_name")
	citizen.MiddleName = form.Get("middle_name")
	citizen.LastName = form.Get("last_name")
	citizen.Gender = form.Get("gender")
	citizen.Major = form.Get("major")
	citizen.TemporaryAddress = form.Get("temporary_address")
	citizen.Avatar = form.Get("avatar")
	citizen.DateOfJoining = form.Get("date_of_joining")
	citizen.DateOfBirth = form.Get("date_of_birth")

	citizen.Religion = form.Get("religion")
	if err != nil {
		errString += "\n" + "weight only contains numbers and must not be empty"
	}
	citizen.CollaboratorName = form.Get("collaborator_name")
	citizen.CollaboratorPhone = form.Get("collaborator_phone")

	dateRegex := `\d{4}-\d{2}-\d{2}`
	codeRegex := "^[0-9]+$"

	if citizen.Gender != "F" && citizen.Gender != "M" {
		errString += "\n" + "gender just only contains 'F' or 'M'"
	}
	if isValid, _ := regexp.MatchString(dateRegex, citizen.DateOfJoining); !isValid {
		errString += "\n" + "Invalid date of joining (must be yyyy-mm-dd)"
	}
	if isValid, _ := regexp.MatchString(dateRegex, citizen.DateOfBirth); !isValid {
		errString += "\n" + "Invalid date of birth (must be yyyy-mm-dd)"
	}
	if isValid, _ := regexp.MatchString(codeRegex, citizen.Code); !isValid {
		errString += "\n" + "CMND/CCCD must have 12 or 9 number"
	} else {
		length := len(citizen.Code)
		if length != 9 && length != 12 {
			errString += "\n" + "CMND/CCCD must have 12 or 9 number"
		}
	}
	if citizen.CollaboratorPhone != "" {
		if isValid, _ := regexp.MatchString(codeRegex, citizen.CollaboratorPhone); !isValid {
			errString += "\n" + "collaborator phone is invalid"
		}
	}
	if citizen.FirstName == "" {
		errString += "\n" + "first name must be not empty"
	}
	if citizen.LastName == "" {
		errString += "\n" + "last name must be not empty"
	}
	if errString == "" {
		err = nil
	} else {
		err = errors.New(errString)
	}
	return citizen, err
}
