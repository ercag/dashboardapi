package main

import (
	"log"
	"net/http"
	"os"

	"dashboardapi/app/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//Path to route
	router.HandleFunc("/", services.Home)
	router.HandleFunc("/login", services.Login)
	router.HandleFunc("/user", services.CreateUser).Methods("POST")

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Println("** Service Started on Port " + port + " **")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
