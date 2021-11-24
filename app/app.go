package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"vietnam-population-server/app/handlers"
	"vietnam-population-server/app/utils"

	_ "github.com/go-sql-driver/mysql"
)

type Handle func(db *sql.DB, w http.ResponseWriter, r *http.Request)

type App struct {
	Router *utils.Router
	server *http.Server
	db     *sql.DB
}

func (a *App) Init() {
	a.Router = utils.NewRouter()
	a.setRouters()
}

func (a *App) Run(host string) {
	a.server = &http.Server{
		Addr:           host,
		Handler:        a.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	var err error
	a.db, err = sql.Open("mysql", "wxKmYNfzWA:uiQirhBvwE@tcp(remotemysql.com:3306)/wxKmYNfzWA")
	if err != nil {
		log.Fatalln("Cannot open mysql")
	}
	log.Printf("Server is running ...")
	log.Fatal(a.server.ListenAndServe())
}

func (a *App) setRouters() {
	a.Router.Add("/", homePage)
	a.Router.Add("/provinces", a.handleRequest(handlers.GetProvinceList))
	a.Router.Add("/districts", a.handleRequest(handlers.GetDistrictListByProvinceCode))
	a.Router.Add("/wards", a.handleRequest(handlers.GetWardListByDistrictCode))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func (a *App) handleRequest(handler Handle) utils.Handle {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.db, w, r)
	}
}
