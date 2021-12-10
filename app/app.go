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

var sqlUrl = "root:root@tcp(localhost:3306)/vietnam_population" // "wxKmYNfzWA:2oVGW6sXGC@tcp(remotemysql.com:3306)/wxKmYNfzWA"

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
	a.db, err = sql.Open("mysql", sqlUrl)
	if err != nil {
		log.Fatalln("Cannot open mysql")
	}
	log.Printf("Server is running ...")

	log.Fatal(a.server.ListenAndServeTLS("server.cert", "server.key"))
}

func (a *App) setRouters() {
	a.Router.Add("/", homePage)
	a.Router.Add("/api/provinces", a.handleRequest(handlers.GetProvinceList))
	a.Router.Add("/api/districts", a.handleRequest(handlers.GetDistrictListByProvinceCode))
	a.Router.Add("/api/wards", a.handleRequest(handlers.GetWardListByDistrictCode))

	a.Router.Add("/api/login", a.handleRequest(handlers.Login))
	a.Router.Add("/api/logout", a.handleRequest(handlers.Logout))

	a.Router.Add("/api/lower-cadres", handlers.IsAuthorized(a.handleRequest(handlers.GetLowerCadreListByCode)))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func (a *App) handleRequest(handler Handle) utils.Handle {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.db, w, r)
	}
}
