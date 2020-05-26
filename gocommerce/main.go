package main

import (
	"gosamples/gocommerce/db"
	"gosamples/gocommerce/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// setup database
	db.SetupDB()

	// Load Routes
	routes.Load()

	// Up server
	http.ListenAndServe(":8000", nil)
}
