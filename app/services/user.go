package services

import (
	"dashboardapi/app/infrastructure/mongodb"
	"dashboardapi/app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person: %+v", user)

	result := mongodb.CreateUser(user)

	fmt.Println("Insert ID: " + result)
}
