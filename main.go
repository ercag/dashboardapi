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

	router.Use(accessControlMiddleware)

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

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" && r.URL.Path != "/login" {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			if services.ValidateToken(token) {
				w.Header().Set("Content-Type", "application/json")
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Token Validation Error. Please control the token.", http.StatusBadRequest)
			}
		}
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		// w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		// if r.Method == "OPTIONS" {
		// 	return
		// }

	})
}
