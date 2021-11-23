package main

import (
	"fmt"
	"log"
	"net/http"
)

var mux = http.NewServeMux()

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	mux.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1000", mux))
}

func main() {
	handleRequests()
}
