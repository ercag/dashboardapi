package services

import (
	"dashboardapi/app/infrastructure/mongodb"
	"dashboardapi/app/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Fprintf(w, "Person: %+v", user)

	result := mongodb.Login(user)

	json.NewEncoder(w).Encode(result)
}

func ValidateToken(token string) (valid bool) {
	return mongodb.ValidateToken(token)
}
