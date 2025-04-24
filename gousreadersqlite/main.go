package main

import (
	"gosamples/gousreadersqlite/db"
	"gosamples/gousreadersqlite/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// setup database
	db.SetupDB()

	// Load Routes
	routes.Load()

	// Up server
	http.ListenAndServe(":8005", nil)
}
