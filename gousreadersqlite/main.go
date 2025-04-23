package main

import (
	"gosamples/gousreadersqlite/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// setup database
	db.SetupDB()
}
