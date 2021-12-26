package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	models "vietnam-population-server/app/models/cadre"
	"vietnam-population-server/app/router"

	"github.com/dgrijalva/jwt-go"
)

var (
	badRequestStatus     = statusCode{number: 400, description: "Bad Request"}
	notImplementedStatus = statusCode{number: 501, description: "Not Implemented"}
	notFoundStatus       = statusCode{number: 404, description: "Not Found"}
	unauthorizedStatus   = statusCode{number: 401, description: "Unauthorized"}
	internalErrorStatus  = statusCode{number: 500, description: "Internal Server Error"}
	mySigningKey         = []byte("this is a line")
	expiredTime          = time.Hour * 4
	blackTokenList       = make(map[string]int)
)

type statusCode struct {
	number      int
	description string
}

type SubdivisionResponse struct {
	AreaSize   float64     `json:"area"`
	Area       string      `json:"area_name"`
	Amount     int         `json:"amount"`
	Population uint32      `json:"population"`
	Data       interface{} `json:"data"`
}

type CitizenListResponse struct {
	Amount int         `json:"amount"`
	Data   interface{} `json:"data"`
}

type CadreListResponse struct {
	Area       string          `json:"area"`
	Permission int             `json:"permission"`
	Amount     int             `json:"amount"`
	Population uint32          `json:"population"`
	Data       []CadreResponse `json:"data"`
}

type CadreResponse struct {
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Age         uint8       `json:"age"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Permission  uint8       `json:"permission"`
	Subdivision interface{} `json:"subdivision"`
}

type JwtResponse struct {
	Token      string `json:"token"`
	Permission int    `json:"permission"`
}

func respondJSON(w *router.ResponseWriter, status int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		respondError(w, internalErrorStatus.number, internalErrorStatus.description)
		return
	}

	header := w.Writer().Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Content-Type", "application/json")
	header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	header.Add("Access-Control-Allow-Headers", "*")
	go w.WriteHeader(status)
	w.Write([]byte(res))
}

func respondError(w *router.ResponseWriter, code int, message string) {
	if strings.HasPrefix(message, "commands out of sync") {
		cmd := exec.Command("service", "vietnam-population", "restart")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	}
	respondJSON(w, code, map[string]string{"error": message})
}

func IsAuthorized(endpoint router.Handle) router.Handle {
	return func(w *router.ResponseWriter, r *http.Request) {
		headerParam := "Authorization"
		if r.Header[headerParam] != nil {
			authString := r.Header[headerParam][0]
			tokenStrings := strings.Split(authString, " ")

			if len(tokenStrings) <= 1 {
				respondError(w, unauthorizedStatus.number, unauthorizedStatus.description)
				return
			}

			tokenString := tokenStrings[1]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(mySigningKey), nil
			})
			if err != nil {
				respondError(w, unauthorizedStatus.number, unauthorizedStatus.description)
				return
			}

			now := time.Now().Add(expiredTime).Unix()
			exp := int64(claims["exp"].(float64))
			isExpired := now-exp > int64(expiredTime)

			if token.Valid && !isExpired && blackTokenList[tokenString] != 1 {
				endpoint(w, r)
			} else {
				respondError(w, unauthorizedStatus.number, unauthorizedStatus.description)
			}
		} else {
			respondError(w, unauthorizedStatus.number, unauthorizedStatus.description)
		}
	}
}

func generateJWT(cadre models.Cadre) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["code"] = cadre.Code
	claims["exp"] = time.Now().Add(expiredTime).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getClaims(r *http.Request) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	headerParam := "Authorization"
	if r.Header[headerParam] != nil {
		authString := r.Header[headerParam][0]
		tokenStrings := strings.Split(authString, " ")
		if len(tokenStrings) <= 1 {
			return claims, errors.New("Missing " + headerParam)
		}

		tokenString := tokenStrings[1]
		claims := jwt.MapClaims{}
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})

		return claims, nil
	}

	return claims, errors.New("Missing " + headerParam)
}

func getPageAndLimit(r *http.Request) (int, int) {
	var (
		defaultPage  = 1
		defaultLimit = 10
	)

	pages, ok1 := r.URL.Query()["page"]
	limits, ok2 := r.URL.Query()["limit"]

	page := defaultPage
	limit := defaultLimit

	if ok1 && ok2 && len(pages) >= 1 || len(limits) >= 1 {
		var err1, err2 error
		page, err1 = strconv.Atoi(pages[0])
		limit, err2 = strconv.Atoi(limits[0])
		if err1 != nil || err2 != nil {
			page = defaultPage
			limit = defaultLimit
		}
	}

	return page, limit
}

func getParam(r *http.Request, key string) (string, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys) < 1 {
		return "", errors.New("URL Param is missing")
	}

	return keys[0], nil
}
