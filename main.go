package main

import (
	"log"
	"net/http"

	"dashboardapi/app/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//Path to route
	router.HandleFunc("/", services.Home)
	router.HandleFunc("/login", services.Login)

	log.Fatal(http.ListenAndServe(":5000", router))
}
