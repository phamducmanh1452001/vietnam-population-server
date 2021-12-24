package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"vietnam-population-server/app/handlers"
	"vietnam-population-server/app/router"

	_ "github.com/go-sql-driver/mysql"
)

type Handle func(db *sql.DB, w *router.ResponseWriter, r *http.Request)

var sqlUrl = "root:root@tcp(localhost:3306)/vietnam_population?multiStatements=true&timeout=5s&tls=false&autocommit=true" // "wxKmYNfzWA:2oVGW6sXGC@tcp(remotemysql.com:3306)/wxKmYNfzWA"

type App struct {
	Router *router.Router
	server *http.Server
	db     *sql.DB
}

func (a *App) Init() {
	a.Router = router.NewRouter()
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
	a.db.SetMaxIdleConns(0)
	a.db.SetMaxOpenConns(500)
	if err != nil {
		log.Fatalln("Cannot open mysql")
	}
	log.Printf("Server is running ...")

	log.Fatal(a.server.ListenAndServe())
}

func (a *App) setRouters() {
	a.Router.Add("/", homePage)
	a.Router.Add("/api/provinces", a.handleRequest(handlers.GetProvinceList))
	a.Router.Add("/api/districts", a.handleRequest(handlers.GetDistrictListByProvinceCode))
	a.Router.Add("/api/wards", a.handleRequest(handlers.GetWardListByDistrictCode))

	a.Router.Add("/api/login", a.handleRequest(handlers.Login))
	a.Router.Add("/api/logout", a.handleRequest(handlers.Logout))

	a.Router.Add("/api/lower-cadres", handlers.IsAuthorized(a.handleRequest(handlers.GetLowerCadreListByCode)))
	a.Router.Add("/api/citizens", handlers.IsAuthorized(a.handleRequest(handlers.GetCitizenList)))
	a.Router.Add("/api/change-cadre-permission", a.handleRequest(handlers.ChangeCadrePermisson))

	a.Router.Add("/api/upload", handlers.UploadImage)
	a.Router.Add("/api/images", handlers.DownloadImage)
}

func homePage(w *router.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func (a *App) handleRequest(handler Handle) func(w *router.ResponseWriter, r *http.Request) {
	return func(w *router.ResponseWriter, r *http.Request) {
		handler(a.db, w, r)
	}
}
