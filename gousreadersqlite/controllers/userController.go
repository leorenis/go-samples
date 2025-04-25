package controllers

import (
	"encoding/json"
	"gosamples/gousreadersqlite/models/users"
	"net/http"
)

// Index is
func Index(w http.ResponseWriter, r *http.Request) {
	users := users.FindAllSuperUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func TopCountries(w http.ResponseWriter, r *http.Request) {
	topFive := users.FindTopFiveCountries()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topFive)
}

func TeamInsights(w http.ResponseWriter, r *http.Request) {
	insights := users.FindInsights()
	json.NewEncoder(w).Encode(insights)
}

// Migrate
func Migrate(w http.ResponseWriter, r *http.Request) {
	users.Migrate()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusAccepted)
}
