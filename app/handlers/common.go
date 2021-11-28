package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	models "vietnam-population-server/app/models/cadre"
	"vietnam-population-server/app/utils"

	"github.com/dgrijalva/jwt-go"
)

var (
	badRequestStatus     = statusCode{number: 400, description: "Bad Request"}
	notImplementedStatus = statusCode{number: 501, description: "Not Implemented"}
	notFoundStatus       = statusCode{number: 404, description: "Not Found"}
	unauthorizedStatus   = statusCode{number: 401, description: "Unauthorized"}
	internalErrorStatus  = statusCode{number: 500, description: "Internal Server Error"}
	mySigningKey         = []byte("this is a line")
	expiredTime          = time.Hour * 1
)

type statusCode struct {
	number      int
	description string
}

type SubdivisionResponse struct {
	Area       string      `json:"area"`
	Population uint32      `json:"population"`
	Data       interface{} `json:"data"`
}

type CadreListResponse struct {
	Area       string          `json:"area"`
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
	Token string `json:"token"`
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res, err := json.Marshal(payload)
	if err != nil {
		respondError(w, internalErrorStatus.number, internalErrorStatus.description)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	respondJSON(w, code, map[string]string{"error": message})
}

func IsAuthorized(endpoint utils.Handle) utils.Handle {
	return func(w http.ResponseWriter, r *http.Request) {
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

			if token.Valid && !isExpired {
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
