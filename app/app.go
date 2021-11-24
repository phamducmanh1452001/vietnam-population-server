package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"vietnam-population-server/app/handlers"
	"vietnam-population-server/app/utils"
)

type App struct {
	Router *utils.Router
	server *http.Server
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
	log.Printf("Server is running ...")
	log.Fatal(a.server.ListenAndServe())
}

func (a *App) setRouters() {
	a.Router.Add("/", homePage)
	a.Router.Add("/provinces", handlers.GetProvinceList)
	a.Router.Add("/districts", handlers.GetDistrictListByProvinceCode)
	a.Router.Add("/wards", handlers.GetWardListByDistrictCode)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}
