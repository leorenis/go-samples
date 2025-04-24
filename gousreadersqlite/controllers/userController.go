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

func Migrate(w http.ResponseWriter, r *http.Request) {
	users.Migrate()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusAccepted)
}
